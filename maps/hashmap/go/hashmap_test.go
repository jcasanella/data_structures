package main

import (
	"data_structures/data/structures/hashmap"
	"strconv"
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

func addElements(hm *hashmap.Hashmap, limit int) {
	for idx := 0; idx < limit; idx++ {
		strIdx := strconv.Itoa(idx)
		hm.Add("key"+strIdx, 10+idx)
	}
}

func checkAllKeys(hm *hashmap.Hashmap, limit int) (string, bool) {
	for idx := 0; idx < limit; idx++ {
		keys := hm.GetKeys()
		strIdx := strconv.Itoa(idx)
		if !contains("key"+strIdx, keys) {
			return "key" + strIdx, true
		}
	}

	return "", false
}

func getAllValues(hm *hashmap.Hashmap, limit int) (int, bool) {
	for idx := 0; idx < limit; idx++ {
		strIdx := strconv.Itoa(idx)
		value, err := hm.Get("key" + strIdx)
		if value != 10+idx || err {
			return value, true
		}
	}

	return 0, false
}

func removeAllElements(hm *hashmap.Hashmap, limit int) (string, bool) {
	for idx := 0; idx < limit; idx++ {
		strIdx := strconv.Itoa(idx)
		if !hm.Remove("key"+strIdx) || int(hm.Size()) != limit-(idx+1) {
			return "key" + strIdx, false
		}

		if _, error := hm.Get("key" + strIdx); error {
			return "key" + strIdx, false
		}
	}

	return "", true
}

func TestMultipleElements(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		key     string
		update1 int
		update2 int
	}{
		{
			name:    "5 elements",
			args:    5,
			key:     "key2",
			update1: 567,
			update2: 378,
		},
		{
			name:    "10 elements",
			args:    10,
			key:     "key7",
			update1: 982,
			update2: 425,
		},
		{
			name:    "40 elements",
			args:    40,
			key:     "key34",
			update1: 643,
			update2: 461,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

			// Add elements to the hashmap
			addElements(hm, tt.args)

			// Get number elements
			if numElems := hm.Size(); numElems != uint(tt.args) {
				t.Errorf("Expected %d elements but actual is %d\n", tt.args, numElems)
			}

			// Get keys
			if keys := hm.GetKeys(); len(keys) != tt.args {
				t.Errorf("Number of keys is not 5, actual is %d\n", len(keys))
			}

			// Check all the keys
			if key, error := checkAllKeys(hm, tt.args); error {
				t.Errorf("Not found key: key%s\n", key)
			}

			// Get All values
			if value, error := getAllValues(hm, tt.args); error {
				t.Errorf("Expected value %d but is different from atual\n", value)
			}

			// Update value
			hm.Add(tt.key, tt.update1)
			if value, _ := hm.Get(tt.key); value != tt.update1 {
				t.Errorf("Value is %d that is different from %d\n", value, tt.update1)
			}
			hm.Add(tt.key, tt.update2)
			if value, _ := hm.Get(tt.key); value != tt.update2 {
				t.Errorf("Value is %d that is different from %d\n", value, tt.update2)
			}

			if numElems := hm.Size(); int(numElems) != tt.args {
				t.Errorf("Expected %d elements but actual is %d\n", tt.args, numElems)
			}

			// Remove all elements
			if value, error := removeAllElements(hm, tt.args); error {
				t.Errorf("Element with key %s not dropped\n", value)
			}
		})
	}
}

// func TestFiveElements(t *testing.T) {
// var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

// addElements(hm, 5)

// Get number elements
// if numElems := hm.Size(); numElems != 5 {
// 	t.Errorf("Expected 5 elements but actual is %d\n", numElems)
// }

// Get keys
// if keys := hm.GetKeys(); len(keys) != 5 {
// t.Errorf("Number of keys is not 5, actual is %d\n", len(keys))
// }

// Check all the keys
// if key, error := checkAllKeys(hm, 5); error {
// t.Errorf("Not found key: key%s\n", key)
// }

// Get All values
// if value, error := getAllValues(hm, 5, 10); error {
// t.Errorf("Expected value %d but is different from atual\n", value)
// }

// // Update value
// hm.Add("key0", 100)
// if value, _ := hm.Get("key0"); value != 100 {
// 	t.Errorf("Value is %d that is different from 100\n", value)
// }
// hm.Add("key0", 10)
// if value, _ := hm.Get("key0"); value != 10 {
// 	t.Errorf("Value is %d that is different from 10\n", value)
// }

// if numElems := hm.Size(); numElems != 5 {
// 	t.Errorf("Expected 5 elements but actual is %d\n", numElems)
// }

// if value, error := removeAllElements(hm, 5); error {
// 	t.Errorf("Element with key %s not dropped\n", value)
// }
// }

/*func TestTenElements(t *testing.T) {
	var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

	addElements(hm, 10)

	// Get number elements
	if numElems := hm.Size(); numElems != 10 {
		t.Errorf("Expected 10 elements but actual is %d\n", numElems)
	}

	// Get keys
	if keys := hm.GetKeys(); len(keys) != 10 {
		t.Errorf("Number of keys is not 10, actual is %d\n", len(keys))
	}

	// Check all the keys
	if key, error := checkAllKeys(hm, 10); error {
		t.Errorf("Not found key: key%s\n", key)
	}

	// Get All values
	if value, error := getAllValues(hm, 10, 10); error {
		t.Errorf("Expected value %d but is different from atual\n", value)
	}

	// Update value
	hm.Add("key8", 123)
	if value, _ := hm.Get("key8"); value != 123 {
		t.Errorf("Value is %d that is different from 123\n", value)
	}
	hm.Add("key8", 567)
	if value, _ := hm.Get("key8"); value != 567 {
		t.Errorf("Value is %d that is different from 567\n", value)
	}

	if numElems := hm.Size(); numElems != 10 {
		t.Errorf("Expected 10 elements but actual is %d\n", numElems)
	}

	if value, error := removeAllElements(hm, 10); error {
		t.Errorf("Element with key %s not dropped\n", value)
	}
}*/
/*
func TestFourtyElements(t *testing.T) {
	var hm *hashmap.Hashmap = hashmap.NewHashmap(10)

	addElements(hm, 40)

	// Get number elements
	if numElems := hm.Size(); numElems != 40 {
		t.Errorf("Expected 40 elements but actual is %d\n", numElems)
	}

	// Get keys
	if keys := hm.GetKeys(); len(keys) != 40 {
		t.Errorf("Number of keys is not 40, actual is %d\n", len(keys))
	}

	// Check all the keys
	if key, error := checkAllKeys(hm, 40); error {
		t.Errorf("Not found key: key%s\n", key)
	}

	// Get All values
	if value, error := getAllValues(hm, 40, 589); error {
		t.Errorf("Expected value %d but is different from atual\n", value)
	}

	// Update value
	hm.Add("key21", 912)
	if value, _ := hm.Get("key21"); value != 912 {
		t.Errorf("Value is %d that is different from 912\n", value)
	}
	hm.Add("key21", 1234)
	if value, _ := hm.Get("key21"); value != 1234 {
		t.Errorf("Value is %d that is different from 1234\n", value)
	}

	if numElems := hm.Size(); numElems != 40 {
		t.Errorf("Expected 40 elements but actual is %d\n", numElems)
	}

	if value, error := removeAllElements(hm, 20); error {
		t.Errorf("Element with key %s not dropped\n", value)
	}
}
*/
func contains(lookFor string, values []string) bool {
	for _, value := range values {
		if value == lookFor {
			return true
		}
	}

	return false
}
