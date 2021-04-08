package main

import (
	"errors"
	"fmt"
)

type ArrayDynamicer interface {
	Get()
	Push()
	Pop()
	DeleteAtIndex()
	ShiftElements()
	GetLength()
	GetCapacity()
}

type Arrays struct {
	data []int
}

func New() Arrays {
	a := Arrays{make([]int, 0, 1)}
	fmt.Printf("len: %d capacity: %d\n", len(a.data), cap(a.data))

	return a
}

func (a Arrays) Get(index int) (int, error) {
	if index < 0 || index > len(a.data) {
		return -1, errors.New("invalid index - can not return value")
	}

	return a.data[index], nil
}

func (a *Arrays) Push(value int) int {
	var newSlice []int
	if len(a.data) == cap(a.data) {
		// We can do something better with append
		newSlice = make([]int, len(a.data)+1, cap(a.data)*2)

	} else {
		newSlice = make([]int, len(a.data)+1, cap(a.data))
	}

	copy(newSlice, a.data)
	a.data = newSlice
	a.data[len(a.data)-1] = value
	fmt.Println(len(a.data), cap(a.data), a.data)

	return value
}

func (a Arrays) Pop() (int, error) {
	if (len(a.data)) > 0 {
		value := a.data[len(a.data)-1]
		a.data = a.data[:len(a.data)-1]

		return value, nil
	}

	return -1, errors.New("empty array - can not pop value")
}

func (a Arrays) GetLength() int {
	return len(a.data)
}

func (a Arrays) GetCapacity() int {
	fmt.Printf("capacity: %d\n", cap(a.data))
	return cap(a.data)
}

func (a Arrays) DeleteAtIndex(index int) (int, error) {
	if index < 0 || index > len(a.data) {
		return -1, errors.New("invalid index - can not delete value")
	}

	value := a.data[index]
	a.ShiftElements(index)

	return value, nil
}

func (a Arrays) ShiftElements(index int) {
	if index >= 0 && index < len(a.data) {
		for idx := index; idx < len(a.data)-1; idx++ {
			a.data[idx] = a.data[idx+1]
		}

		a.data = a.data[:len(a.data)-1]
	}
}
