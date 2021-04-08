package main

import (
	"testing"
)

func ValidateLengthCapacity(t *testing.T, actualLen int, expectedLen int, actualCap int, expectedCap int) {
	if actualLen != expectedLen {
		t.Errorf("Length is incorrect, got %d, want %d", actualLen, expectedLen)
	}

	if actualCap != expectedCap {
		t.Errorf("Capacity is incorrect, got %d want %d", actualCap, expectedCap)
	}
}

func ValidateValue(t *testing.T, actualValue int, expectedValue int) {
	if actualValue != expectedValue {
		t.Errorf("Value is incorrect, got %d want %d", actualValue, expectedValue)
	}
}

func TestEmptyArray(t *testing.T) {
	var dynamic Arrays = New()

	ValidateLengthCapacity(t, dynamic.GetLength(), 0, dynamic.GetCapacity(), 1)
}

func TestPush(t *testing.T) {
	var dynamic Arrays = New()

	value := dynamic.Push(1)
	ValidateLengthCapacity(t, dynamic.GetLength(), 1, dynamic.GetCapacity(), 1)
	ValidateValue(t, value, 1)

	value = dynamic.Push(2)
	ValidateLengthCapacity(t, dynamic.GetLength(), 2, dynamic.GetCapacity(), 2)
	ValidateValue(t, value, 2)

	value = dynamic.Push(3)
	ValidateLengthCapacity(t, dynamic.GetLength(), 3, dynamic.GetCapacity(), 4)
	ValidateValue(t, value, 3)

	value = dynamic.Push(4)
	ValidateLengthCapacity(t, dynamic.GetLength(), 4, dynamic.GetCapacity(), 4)
	ValidateValue(t, value, 4)

	dynamic.Push(5)
	dynamic.Push(6)
	value = dynamic.Push(7)
	ValidateLengthCapacity(t, dynamic.GetLength(), 7, dynamic.GetCapacity(), 8)
	ValidateValue(t, value, 7)
}
