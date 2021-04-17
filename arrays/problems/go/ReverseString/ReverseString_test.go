package main

import "testing"

func TestEmptyString(t *testing.T) {
	actual := ReverseString("")
	if actual != "" {
		t.Errorf("Actual %s is different than expected %s", actual, "")
	}

	if len(actual) != 0 {
		t.Errorf("Length %d is different to %d", len(actual), 0)
	}
}

func TestStringOneChar(t *testing.T) {
	expected := "a"
	actual := ReverseString(expected)

	if actual != expected {
		t.Errorf("Actual %s is different than expected %s", actual, expected)
	}

	if len(actual) != len(expected) {
		t.Errorf("Length %d is different to %d", len(actual), len(expected))
	}
}

func TestStringOneWord(t *testing.T) {
	word := "skate"
	expected := "etaks"
	actual := ReverseString(word)

	if actual != expected {
		t.Errorf("Actual %s is different than expected %s", actual, expected)
	}

	if len(actual) != len(expected) {
		t.Errorf("Length %d is different to %d", len(actual), len(expected))
	}
}

func TestStringOneWordWithSpaces(t *testing.T) {
	word := " skate "
	expected := " etaks "
	actual := ReverseString(word)

	if actual != expected {
		t.Errorf("Actual %s is different than expected %s", actual, expected)
	}

	if len(actual) != len(expected) {
		t.Errorf("Length %d is different to %d", len(actual), len(expected))
	}
}

func TestStringMultipleWords(t *testing.T) {
	word := "Inline skates is cool"
	expected := "looc si setaks enilnI"
	actual := ReverseString(word)

	if actual != expected {
		t.Errorf("Actual %s is different than expected %s", actual, expected)
	}

	if len(actual) != len(expected) {
		t.Errorf("Length %d is different to %d", len(actual), len(expected))
	}
}
