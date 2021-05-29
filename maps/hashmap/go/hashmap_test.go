package main

import (
	"data_structures/data/structures/hashmap"
	"testing"

)

// Add(key string, value int)
// Remove(key string) bool
// Get(key string) (int, bool)
// GetKeys() []string
// Size() uint

func TestEmpty(t *testing.T) {
	var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

	// Must be empty
	if numElems := hm.Size(); numElems != 0 {
		t.Errorf("Expected size is %d but actual is 0\n", numElems)
	}

	// Get Keys
	if keys := hm.GetKeys(); len(keys) != 0 {
		t.Errorf("Expected size is %d but actual is 0\n", len(keys))
	}

	// Get element
	if value, error := hm.Get("key1"); !error {
		t.Errorf("Expected an error, but no error and returned value %d\n", value)
	}

	if removed := hm.Remove("key1"); removed {
		t.Errorf("Removed when no values in the hashMap")
	}
}

func TestAddElem(t *testing.T) {
	var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

	hm.Add("key1", 10)
	if numElems := hm.Size(); numElems != 1 {
		t.Errorf("Expected size 1 but is %d\n", numElems)
	}

	// Must be one
	if numElems := hm.Size(); numElems != 1 {
		t.Errorf("Expected size is %d but actual is 1\n", numElems)
	}

	// Get element
	if value, error := hm.Get("key1"); value != 10 || error {
		t.Errorf("Expected value 10 but get %d and noError but was %t\n", value, error)
	}

	// Get keys
	if keys := hm.GetKeys(); len(keys) != 1 {
		t.Errorf("Expected 1 key but got %d\n", len(keys))
	}
	if keys := hm.GetKeys(); keys[0] != "key1" {
		t.Errorf("Expected key1 key but got %v\n", keys[0])
	}

	// Overwrite value
	hm.Add("key1", 15)
	if numElems := hm.Size(); numElems != 1 {
		t.Errorf("Expected size 1 but is %d\n", numElems)
	}

	// Must be one
	if numElems := hm.Size(); numElems != 1 {
		t.Errorf("Expected size is %d but actual is 1\n", numElems)
	}

	// Get element
	if value, error := hm.Get("key1"); value != 15 || error {
		t.Errorf("Expected value 15 but get %d and noError but was %t\n", value, error)
	}

	// Get keys
	if keys := hm.GetKeys(); len(keys) != 1 {
		t.Errorf("Expected 1 key but got %d\n", len(keys))
	}
	if keys := hm.GetKeys(); keys[0] != "key1" {
		t.Errorf("Expected key1 key but got %v\n", keys[0])
	}

	// Remove Element
	if removed := hm.Remove("key1"); !removed {
		t.Errorf("Value not removed\n")
	}

	// Must be empty
	if numElems := hm.Size(); numElems != 0 {
		t.Errorf("Expected size is %d but actual is 0\n", numElems)
	}

	// Get keys - empty
	if keys := hm.GetKeys(); len(keys) != 0 {
		t.Errorf("Expected 0 key but got %d\n", len(keys))
	}

	// Get element
	if value, error := hm.Get("key1"); !error {
		t.Errorf("Expected error but got noError and value %d\n", value)
	}
}
