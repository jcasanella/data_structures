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
	numElems := hm.Size()
	if numElems != 0 {
		t.Errorf("Expected size is %d but actual is 0\n", numElems)
	}

	// Get Keys
	keys := hm.GetKeys()
	if len(keys) != 0 {
		t.Errorf("Expected size is %d but actual is 0\n", numElems)
	}

	// Get element
	if value, error := hm.Get("key1"); !error {
		t.Errorf("Expected an error, but no error and returned value %d\n", value)
	}

	if removed := hm.Remove("key1"); removed {
		t.Errorf("Removed when no values in the hashMap")
	}
}
