package main

import (
	"orderedlist/src/constraints"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	sut := OrderedList[int]{}

	assert.Equal(t, 0, sut.Count())

}

func TestAddAscTwoElems(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	sut.Add(1)
	sut.Add(2)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.head, sut.tail.prev)
	checkAsc(sut, t)
}

func TestAddDescTwoElems(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	sut.Add(1)
	sut.Add(2)

	assert.Equal(t, sut.head.next, sut.tail)
	assert.Equal(t, sut.head, sut.tail.prev)
	assert.Equal(t, sut.head.value, 2)
	assert.Equal(t, sut.tail.value, 1)
	checkDesc(sut, t)
}

func TestAddAscInEmpty(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	sut.Add(1)

	assert.Equal(t, sut.head, sut.tail)
	assert.Equal(t, sut.head.value, 1)
	checkAsc(sut, t)
}

func TestAddAsc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8, 0}

	for _, item := range items {
		sut.Add(item)
	}

	assert.Equal(t, len(items), sut.Count())
	assert.Equal(t, 8, sut.tail.value)
	checkAsc(sut, t)
}

func FuzzAddIntAsc(f *testing.F) {
	sut := OrderedList[int]{nil, nil, true}
	f.Fuzz(func(t *testing.T, num int) {
		sut.Add(num)

		checkAsc[int](sut, t)
	})
}

func FuzzAddStringAsc(f *testing.F) {
	sut := OrderedList[string]{nil, nil, true}
	f.Fuzz(func(t *testing.T, num string) {
		sut.Add(num)

		checkAsc(sut, t)
	})
}

func TestAddDesc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8, 0}

	for _, item := range items {
		sut.Add(item)
	}

	assert.Equal(t, len(items), sut.Count())
	assert.Equal(t, -11, sut.tail.value)
	checkDesc(sut, t)
}

func TestFindInEmpty(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}

	_, err := sut.Find(-1)

	assert.NotNil(t, err)
}

func TestFindExistingAsc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8, 0}

	for _, item := range items {
		sut.Add(item)
	}
	checkAsc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(item)

			assert.Nil(t, err)
			assert.Equal(t, item, res.value)
		})
	}

	assert.Equal(t, sut.tail.value, 8)
	assert.Equal(t, sut.head.value, -11)
}

func TestFindExistingDesc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8, 0}

	for _, item := range items {
		sut.Add(item)
	}
	checkDesc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(item)

			assert.Nil(t, err)
			assert.Equal(t, item, res.value)
		})
	}
}

func TestFindNotExistingAsc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8}

	for _, item := range items {
		sut.Add(item)
	}
	checkAsc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			_, err := sut.Find(-1 * item)

			assert.NotNil(t, err)
		})
	}
}

func TestFindNotExistingDesc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8}

	for _, item := range items {
		sut.Add(item)
	}
	checkDesc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			_, err := sut.Find(-1 * item)

			assert.NotNil(t, err)
		})
	}
}

func TestDeleteEmpty(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	old := sut.Count()
	sut.Delete(1)
	new := sut.Count()

	assert.Equal(t, old, new)
}

func TestDeleteExistingDesc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, false}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8}

	for _, item := range items {
		sut.Add(item)
	}
	checkDesc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			old := sut.Count()
			sut.Delete(item)
			new := sut.Count()

			assert.Equal(t, old-1, new)
			checkDesc(sut, t)

		})
	}

	assert.Nil(t, sut.head)
}

func TestDeleteExistingAsc(t *testing.T) {
	sut := OrderedList[int]{nil, nil, true}
	items := []int{1, 1, 2, 3, 4, 5, 5, 6, -11, 6, 7, 8}

	for _, item := range items {
		sut.Add(item)
	}
	checkAsc(sut, t)

	for _, item := range items {
		t.Run("", func(t *testing.T) {
			old := sut.Count()
			sut.Delete(item)
			new := sut.Count()

			assert.Equal(t, old-1, new)
			checkAsc(sut, t)

		})
	}

	assert.Nil(t, sut.head)
}

func FuzzAddIntDesc(f *testing.F) {
	sut := OrderedList[int]{nil, nil, false}
	f.Fuzz(func(t *testing.T, num int) {
		sut.Add(num)

		checkDesc[int](sut, t)
	})
}

func FuzzAddStringDesc(f *testing.F) {
	sut := OrderedList[string]{nil, nil, false}
	f.Fuzz(func(t *testing.T, num string) {
		sut.Add(num)

		checkDesc(sut, t)
	})
}

func checkAsc[T constraints.Ordered](l OrderedList[T], t *testing.T) {
	tmp := l.head
	if tmp == nil || tmp.next == nil {
		assert.True(t, true)
		return
	}
	for tmp != nil && tmp.next != nil {
		assert.LessOrEqual(t, tmp.value, tmp.next.value)
		if tmp != l.head {
			assert.NotNil(t, tmp.prev)
		}
		assert.NotNil(t, tmp.next)
		tmp = tmp.next
	}
}

func checkDesc[T constraints.Ordered](l OrderedList[T], t *testing.T) {
	tmp := l.head
	if tmp == nil || tmp.next == nil {
		assert.True(t, true)
		return
	}
	for tmp != nil && tmp.next != nil {
		assert.GreaterOrEqual(t, tmp.value, tmp.next.value)
		if tmp != l.head {
			assert.NotNil(t, tmp.prev)
		}
		assert.NotNil(t, tmp.next)
		tmp = tmp.next
	}
}
