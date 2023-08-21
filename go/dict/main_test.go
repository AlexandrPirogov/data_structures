package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzCorrectHashFun(f *testing.F) {
	sz := 16
	sut := Init[string](sz)

	f.Fuzz(func(t *testing.T, s string) {
		index := sut.HashFun(s)

		assert.GreaterOrEqual(t, index, 0)
		assert.Less(t, index, sz)
	})
}

func TestPutInEmpty(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, val := "alex", "pirogov"
	index := sut.HashFun(k)
	sut.Put(k, val)

	assert.Equal(t, sut.slots[index], k)
	assert.Equal(t, sut.values[index], val)
}

func TestGetFromEmpty(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	_, err := sut.Get("not existing")

	assert.NotNil(t, err)
}

func TestGetExistingWithoutCollision(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, val := "alex", "pirogov"
	sut.Put(k, val)
	res, err := sut.Get("alex")

	assert.Nil(t, err)
	assert.Equal(t, val, res)
}

func TestInKeyInEmpty(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	b := sut.IsKey("not existing")

	assert.False(t, b)
}

func TestInKeyNotExistingInFilled(t *testing.T) {
	sz := 3
	sut := Init[string](sz)
	keys := []string{"a", "b", "c"}
	for _, k := range keys {
		sut.Put(k, "asd")
	}
	b := sut.IsKey("not existing")

	assert.False(t, b)
}

func TestInKeyExistingInFilled(t *testing.T) {
	sz := 3
	sut := Init[string](sz)
	keys := []string{"a", "b", "c"}
	for _, k := range keys {
		sut.Put(k, "asd")
	}

	for _, k := range keys {
		t.Run(k, func(t *testing.T) {
			b := sut.IsKey(k)
			assert.True(t, b)
		})
	}

}

func TestInKeyExistingWithoutCollision(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, val := "alex", "pirogov"
	sut.Put(k, val)
	b := sut.IsKey("alex")

	assert.True(t, b)
}

func TestRewrite(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, old, new := "alex", "pirogov", "awesome"
	sut.Put(k, old)
	sut.Put(k, new)

	res, err := sut.Get("alex")

	assert.Nil(t, err)
	assert.Equal(t, new, res)
}

func TestPutWithCollision(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, val := "alex", "pirogov"
	k1, val1 := "lexa", "ripogov"

	sut.Put(k, val)
	sut.Put(k1, val1)

	res, err := sut.Get(k)
	res1, err1 := sut.Get(k1)

	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Equal(t, val, res)
	assert.Equal(t, val1, res1)

}

func TestRewriteWithCollision(t *testing.T) {
	sz := 16
	sut := Init[string](sz)

	k, old := "alex", "pirogov"
	k1, new := "lexa", "ripogov"

	sut.Put(k, old)
	sut.Put(k1, new)
	sut.Put(k1, old)

	res, err := sut.Get(k1)

	assert.Nil(t, err)
	assert.Equal(t, res, old)
}

func TestFuzzCustom(t *testing.T) {
	sz := 7
	sut := Init[string](sz)

	k, v := "0", "0"
	sut.Put(k, v)

	res, err := sut.Get(k)

	assert.Nil(t, err)
	assert.Equal(t, res, v)
}
