package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := Init[int](16)

	assert.Equal(t, 16, sut.size)
}

func FuzzHash(f *testing.F) {
	sut := Init[string](16)

	f.Fuzz(func(t *testing.T, str string) {
		index := sut.hash(str)
		assert.GreaterOrEqual(t, index, 0)
		assert.Less(t, index, len(sut.slots))
	})

}

func TestAdd(t *testing.T) {
	sut := Init[string](16)

	k, v := "key", "value"
	sut.Put(k, v)
	res, err := sut.Get(k)

	assert.Nil(t, err)
	assert.Equal(t, v, res)
}

func TestAddFull(t *testing.T) {
	sut := Init[string](5)

	strings := []string{"a", "b", "c", "d", "e"}
	for _, v := range strings {
		sut.Put(v, v)
	}

	for _, test := range strings {
		t.Run(test, func(t *testing.T) {
			res, err := sut.Get(test)

			assert.Nil(t, err)
			assert.Equal(t, test, res)
		})
	}
}

func TestAddFullHits(t *testing.T) {
	sut := Init[string](5)

	strings := []string{"a", "b", "c", "d", "e"}
	for _, v := range strings {
		sut.Put(v, v)
	}

	for _, test := range sut.hits {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, 1, test)
		})
	}
}

func TestAddHitsIncrease(t *testing.T) {
	sut := Init[string](5)

	strings := []string{"a", "b", "c", "d", "e"}
	for _, v := range strings {
		sut.Put(v, v)
		sut.Put(v, v)
	}

	for _, test := range sut.hits {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, 2, test)
		})
	}
}

func TestAddDelete(t *testing.T) {
	sut := Init[string](5)

	strings := []string{"a", "b", "c", "d", "e"}
	for _, v := range strings {
		sut.Put(v, v)
		sut.Put(v, v)
	}

	sut.Put("h", "h")
	res, err := sut.Get("h")

	assert.Nil(t, err)
	assert.Equal(t, res, "h")
}

func TestAddCollision(t *testing.T) {
	sut := Init[string](5)

	strings := []string{"ae", "ea"}
	for _, v := range strings {
		sut.Put(v, v)
	}

	for _, test := range strings {
		t.Run(test, func(t *testing.T) {
			res, err := sut.Get(test)

			assert.Nil(t, err)
			assert.Equal(t, test, res)
		})
	}
}
