package main

import (
	"data_structures/data/structures/hashmap"
	"fmt"

)

func main() {
	// Follow these steps to enable mod
	// go mod init github.com/jcasanella/data_structures
	//	go env -w GO111MODULE=on
	// go env -w GO111MODULE=off
	var hm *hashmap.Hashmap = hashmap.NewHashmap(3)
	hm.Add("k", 2)
	hm.Add("a", 3)
	hm.Add("b", 4)
	hm.Add("c", 5)
	hm.Add("d", 6)
	hm.Add("e", 7)
	hm.Add("f", 8)

	numElems := hm.Size()
	fmt.Printf("%d elements \n", numElems)

	if value1, error := hm.Get("k"); !error {
		fmt.Println("Key: k", value1)
	}

	if value1, error := hm.Get("a"); !error {
		fmt.Println("Key: a", value1)
	}

	if value1, error := hm.Get("b"); !error {
		fmt.Println("Key: b", value1)
	}

	if value1, error := hm.Get("c"); !error {
		fmt.Println("Key: c", value1)
	}

	if value1, error := hm.Get("d"); !error {
		fmt.Println("Key: d", value1)
	}

	if value1, error := hm.Get("e"); !error {
		fmt.Println("Key: e", value1)
	}

	if value1, error := hm.Get("f"); !error {
		fmt.Println("Key: f", value1)
	}
}
