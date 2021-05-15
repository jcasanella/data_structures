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
	var hm *hashmap.Hashmap = hashmap.NewHashmap(5)
	hm.Add("k", 5)

	numElems := hm.Size()
	fmt.Printf("%d elements \n", numElems)
}
