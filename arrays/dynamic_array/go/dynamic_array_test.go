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

func ValidateValueAndError(t *testing.T, actualValue int, expectedValue int, actualError error) {
	if actualValue != expectedValue {
		t.Errorf("Value is incorrect, got %d want %d", actualValue, expectedValue)
	}

	if actualError != nil {
		t.Errorf("Error unexpedted error = %v", actualError)
	}
}

func TestGet(t *testing.T) {
	var dynamic Arrays = New()

	_, error := dynamic.Get(-1)
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}

	dynamic.Push(1)
	value, error := dynamic.Get(0)
	ValidateValueAndError(t, value, 1, error)
	ValidateLengthCapacity(t, dynamic.GetLength(), 1, dynamic.GetCapacity(), 1)

	dynamic.Push(2)
	dynamic.Push(34)
	dynamic.Push(70)
	dynamic.Push(45)

	value, error = dynamic.Get(4)
	ValidateValueAndError(t, value, 45, error)
	ValidateLengthCapacity(t, dynamic.GetLength(), 5, dynamic.GetCapacity(), 8)

	value, error = dynamic.Get(3)
	ValidateValueAndError(t, value, 70, error)

	value, error = dynamic.Get(2)
	ValidateValueAndError(t, value, 34, error)

	value, error = dynamic.Get(1)
	ValidateValueAndError(t, value, 2, error)

	_, error = dynamic.Get(19)
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}
}

func TestPop(t *testing.T) {
	var dynamic Arrays = New()

	_, error := dynamic.Pop()
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}

	actual := dynamic.Push(50)
	expected, _ := dynamic.Pop()
	if actual != expected {
		t.Errorf("Error Actual %d and Expected %d are different", actual, expected)
	}
	ValidateLengthCapacity(t, dynamic.GetLength(), 0, dynamic.GetCapacity(), 1)

	dynamic.Push(2)
	dynamic.Push(4)
	dynamic.Push(8)
	dynamic.Push(16)
	dynamic.Push(32)
	dynamic.Push(64)
	dynamic.Push(128)
	ValidateLengthCapacity(t, dynamic.GetLength(), 7, dynamic.GetCapacity(), 8)

	values := [...]int{128, 64, 32, 16, 8, 4, 2}
	for _, s := range values {
		actual, _ = dynamic.Pop()
		if actual != s {
			t.Errorf("Error Actual %d and Expected %d are different", actual, s)
		}
	}

	_, error = dynamic.Pop()
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}
}

func TestDeleteAtIndex(t *testing.T) {
	var dynamic Arrays = New()

	_, error := dynamic.DeleteAtIndex(0)
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}

	dynamic.Push(2)
	dynamic.Push(4)
	dynamic.Push(8)
	dynamic.Push(16)
	dynamic.Push(32)
	dynamic.Push(64)
	dynamic.Push(128)

	ValidateLengthCapacity(t, dynamic.GetLength(), 7, dynamic.GetCapacity(), 8)

	actual, error := dynamic.DeleteAtIndex(2)
	ValidateValueAndError(t, actual, 8, error)
	ValidateLengthCapacity(t, dynamic.GetLength(), 6, dynamic.GetCapacity(), 8)

	actual, error = dynamic.Get(2)
	ValidateValueAndError(t, actual, 16, error)

	actual, error = dynamic.DeleteAtIndex(0)
	ValidateValueAndError(t, actual, 2, error)
	ValidateLengthCapacity(t, dynamic.GetLength(), 5, dynamic.GetCapacity(), 8)

	actual, error = dynamic.Get(0)
	ValidateValueAndError(t, actual, 4, error)

	actual, error = dynamic.DeleteAtIndex(3)
	ValidateValueAndError(t, actual, 64, error)
	ValidateLengthCapacity(t, dynamic.GetLength(), 4, dynamic.GetCapacity(), 8)

	actual, error = dynamic.Get(3)
	ValidateValueAndError(t, actual, 128, error)

	for idx := 3; idx >= 0; idx-- {
		dynamic.DeleteAtIndex(idx)
		ValidateLengthCapacity(t, dynamic.GetLength(), idx, dynamic.GetCapacity(), 8)
	}

	_, error = dynamic.DeleteAtIndex(0)
	if error == nil {
		t.Errorf("Error actual = %v", error)
	}
}
