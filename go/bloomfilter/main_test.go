package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strings []string = []string{
	"0123456789",
	"1234567890",
	"2345678901",
	"3456789012",
	"4567890123",
	"5678901234",
	"6789012345",
	"7890123456",
	"8901234567",
	"9012345678",
}

func TestAddInEmpty(t *testing.T) {
	sut := BloomFilter{filter_len: 32}

	for _, v := range strings {
		t.Run(v, func(t *testing.T) {
			assert.False(t, sut.IsValue(v))
		})
	}
}

func TestAddWithSingle(t *testing.T) {
	sut := BloomFilter{filter_len: 32}

	sut.Add(strings[0])

	assert.True(t, sut.IsValue(strings[0]))
}

func TestAddPrettyEmpty(t *testing.T) {
	sut := BloomFilter{filter_len: 32}
	for i := 0; i < len(strings)-1; i++ {
		sut.Add(strings[0])
	}

	assert.False(t, sut.IsValue(strings[9]))
}

func TestAddAndExist(t *testing.T) {
	sut := BloomFilter{filter_len: 32}
	for _, v := range strings {
		sut.Add(v)
	}

	for _, v := range strings {
		t.Run(v, func(t *testing.T) {
			assert.True(t, sut.IsValue(v))
		})
	}

}
