package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := Init[int]()

	assert.Empty(t, sut.slots)
}

func TestPutInEmpty(t *testing.T) {
	sut := Init[int]()

	sut.Put(1)

	assert.Equal(t, sut.Size(), 1)
}

func TestPutTwitchSame(t *testing.T) {
	sut := Init[int]()

	sut.Put(1)
	sut.Put(1)
	assert.Equal(t, sut.Size(), 1)

}

func TestGetNotExisting(t *testing.T) {
	sut := Init[int]()

	r := sut.Get(-1)

	assert.False(t, r)
}

func TestGetExistingInFilled(t *testing.T) {
	sut := Init[int]()

	sut.Put(1)
	r := sut.Get(1)

	assert.True(t, r)
}

func TestRemoveFromEmpty(t *testing.T) {
	sut := Init[int]()

	r := sut.Remove(1)

	assert.False(t, r)
}

func TestRemoveExisting(t *testing.T) {
	sut := Init[int]()

	sut.Put(1)
	r := sut.Remove(1)

	assert.True(t, r)
}

func TestIntersectionWithoutIntersect(t *testing.T) {
	sut1 := Init[int]()
	sut2 := Init[int]()

	for i := 0; i < 5; i++ {
		sut1.Put(i - 10)
		sut2.Put(i + 10)
	}

	r1 := sut1.Intersection(sut2)
	r2 := sut2.Intersection(sut1)

	assert.Equal(t, r1, r2)

}

func TestIntersectionWithIntersect(t *testing.T) {
	sut1 := Init[int]()
	sut2 := Init[int]()

	for i := 0; i < 5; i++ {
		sut1.Put(i)
		sut2.Put(i)
	}

	r1 := sut1.Intersection(sut2)
	r2 := sut2.Intersection(sut1)

	assert.NotEmpty(t, r1)
	assert.NotEmpty(t, r2)
	assert.Equal(t, r1, r2)

}

func TestUnion(t *testing.T) {
	sut1 := Init[int]()
	sut2 := Init[int]()

	for i := 0; i < 5; i++ {
		sut1.Put(i - 10)
		sut2.Put(i + 10)
	}

	r1 := sut1.Union(sut2)

	assert.NotEmpty(t, r1)
	assert.Equal(t, sut1.Size()+sut2.Size(), r1.Size())
}

func TestDifference(t *testing.T) {
	sut1 := Init[int]()
	sut2 := Init[int]()

	for i, k := 0, 3; i < 5; i, k = i+1, k+1 {
		sut1.Put(i)
		sut2.Put(k)
	}

	r1 := sut1.Difference(sut2)
	r2 := sut2.Difference(sut1)

	assert.Equal(t, 3, r1.Size())
	assert.Equal(t, 3, r2.Size())

}

func TestIsSubset(t *testing.T) {
	sut1 := Init[int]()
	sut2 := Init[int]()

	for i := 0; i < 5; i++ {
		sut1.Put(i)
	}

	for i := 2; i < 4; i++ {
		sut2.Put(i)
	}

	assert.True(t, sut1.IsSubset(sut2))

}
