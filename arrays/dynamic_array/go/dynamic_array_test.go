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

/*
public void testDeleteAtIndex() {
	Integer value = ad.deleteAtIndex(0);
	assertNull(value, "Should be null");

	ad.push(2);
	ad.push(4);
	ad.push(8);
	ad.push(16);
	ad.push(32);
	ad.push(64);
	ad.push(128);
	assertEquals(8, ad.getCapacity(), "Capacity should be 8");
	assertEquals(7, ad.getLength(), "Length should be 7");

	value = ad.get(2);
	Integer value2 = ad.deleteAtIndex(2);
	assertEquals(value, value2, "Should be the same value");
	assertEquals(6, ad.getLength(), "Length should be 6");
	assertEquals(8, ad.getCapacity(), "Capacity should be 8");
	value = ad.get(2);
	assertEquals(16, value, "Should be the same value");

	value = ad.get(0);
	value2 = ad.deleteAtIndex(0);
	assertEquals(value, value2, "Should be the same value");
	assertEquals(5, ad.getLength(), "Length should be 5");
	assertEquals(8, ad.getCapacity(), "Capacity should be 8");
	value = ad.get(0);
	assertEquals(4, value, "Should be the same value");

	value = ad.get(3);
	value2 = ad.deleteAtIndex(3);
	assertEquals(value, value2, "Should be the same value");
	assertEquals(4, ad.getLength(), "Length should be 4");
	assertEquals(8, ad.getCapacity(), "Capacity should be 8");
	value = ad.get(3);
	assertEquals(128, value, "Should be the same value");

	ad.deleteAtIndex(3);
	assertEquals(3, ad.getLength(), "Length should be 3");
	ad.deleteAtIndex(2);
	assertEquals(2, ad.getLength(), "Length should be 2");
	ad.deleteAtIndex(1);
	assertEquals(1, ad.getLength(), "Length should be 1");
	ad.deleteAtIndex(0);
	assertEquals(0, ad.getLength(), "Length should be 0");

	value = ad.deleteAtIndex(50);
	assertNull(value, "Should be a null value");
}
*/
