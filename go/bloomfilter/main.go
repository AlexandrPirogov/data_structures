package main

import (
	"os"
)

var bits int = 32

// битовый массив длиной f_len ...
type BloomFilter struct {
	filter_len int
}

// хэш-функции
// 17
func (b *BloomFilter) Hash1(s string) int {
	sum := 0
	prev := 0
	for _, char := range s {
		code := int(char)
		iter := ((code + prev*17) % b.filter_len)
		prev = iter
		sum = iter
	}
	// реализация ...
	return sum
}

// 223
func (b *BloomFilter) Hash2(s string) int {
	sum := 0
	prev := 0
	for _, char := range s {
		code := int(char)
		iter := ((code + prev*223) % b.filter_len)
		prev = iter
		sum = iter
	}
	// реализация ...
	return sum
}

// добавляем строку s в фильтр
func (b *BloomFilter) Add(s string) {
	n1, n2 := b.Hash1(s), b.Hash2(s)
	bits = (bits | n1) | (bits | n2)
}

// проверка, имеется ли строка s в фильтре
func (b *BloomFilter) IsValue(s string) bool {
	n1, n2 := b.Hash1(s), b.Hash2(s)
	return (bits&n1) == n1 && (bits&n2) == n2
}
