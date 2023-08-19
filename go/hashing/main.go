package main

type HashTable struct {
	size  int
	step  int
	slots []string
	busy  []bool
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	ht.busy = make([]bool, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	sum := 0
	for _, char := range value {
		sum += int(char)
	}
	return sum % ht.size
}

// Evaluate ind for value and return -1 if there are no empty slots
// Otherwise returns index of empty slot
func (ht *HashTable) SeekSlot(value string) int {
	ind := ht.HashFun(value)
	if !ht.busy[ind] {
		return ind
	}

	// находит индекс слота со значением, или -1
	return ht.seek(ind)
}

// записываем значение по хэш-функции
// возвращается индекс слота или -1
// если из-за коллизий элемент не удаётся разместить
func (ht *HashTable) Put(value string) int {
	ind := ht.HashFun(value)
	if ht.busy[ind] {
		return -1

	}

	ht.slots[ind] = value
	ht.busy[ind] = true

	return ind
}

// Checks if given value existing in slots
func (ht *HashTable) Find(value string) int {
	for k, str := range ht.slots {
		if str == value {
			return k
		}
	}
	// находит индекс слота со значением, или -1
	return -1
}

func (ht *HashTable) seek(from int) int {
	if ht.filled() {
		return -1
	}

	free := from
	for ht.busy[free] {
		free = (free + ht.step) % ht.size
	}
	return free
}

func (ht *HashTable) filled() bool {
	for _, b := range ht.busy {
		if !b {
			return false
		}
	}
	return true
}
