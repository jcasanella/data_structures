package hashmap

import "fmt"

type bucket struct {
	key   string
	value int
	next  *bucket
}

type Hashmap struct {
	buckets     []bucket
	numElements uint
}

type Hashmaper interface {
	Add(key string, value int)
	Remove(key string) bool
	Get(key string) (int, bool)
	GetKeys() []string
	Size() uint
}

func NewHashmap(numBuckets int) *Hashmap {
	hashmap := new(Hashmap)
	hashmap.buckets = make([]bucket, numBuckets)
	hashmap.numElements = 0
	return hashmap
}

func (fm *Hashmap) getBucketIndex(key string) int {
	hash := 0
	r := []rune(key)
	for _, value := range r {
		hash = (13*hash + int(value)) % len(fm.buckets)
	}

	return hash
}

func (fm *Hashmap) addWithNoCollision(key string, value int, posic int) {
	fm.buckets[posic].key = key
	fm.buckets[posic].value = value
	fm.numElements++
}

// func (fm *Hashmap) addWithCollision(key string, value int, posic int) {

// }

func (hm *Hashmap) Add(key string, value int) {
	posic := hm.getBucketIndex(key)

	if hm.buckets[posic].key == "" {
		hm.addWithNoCollision(key, value, posic)
	}
	// } else {
	// 	addWithCollision(key, value, posic)
	// }
}

func (hm *Hashmap) Size() uint {
	return hm.numElements
}

func (hm *Hashmap) Get(key string) (int, bool) {
	posic := hm.getBucketIndex(key)
	if hm.buckets[posic].key == "" {
		return 0, true
	}

	_ = posic

	fmt.Printf("%v\n", hm.buckets)

	//if hm.buckets[posic] == nil {
	//
	//	}

	return 0, false
}
