package hashmap

import "fmt"

type bucket struct {
	key   string
	value int
	next  *bucket
}

type Hashmap struct {
	buckets     []*bucket
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
	hashmap.buckets = make([]*bucket, numBuckets)
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
	newBucket := new(bucket)
	newBucket.key = key
	newBucket.value = value

	fm.buckets[posic] = newBucket
	fm.numElements++
}

func (fm *Hashmap) addWithCollision(key string, value int, posic int) {
	existingBucket := fm.buckets[posic]

	// Head element
	if existingBucket.key == key {
		existingBucket.value = value
		return
	}

	// Check if exists in the linked list
	for existingBucket.next != nil {
		existingBucket = existingBucket.next
		if existingBucket.key == key {
			existingBucket.value = value
			return
		}
	}

	// It's a new element
	newNode := bucket{}
	newNode.key = key
	newNode.value = value
	existingBucket.next = &newNode
	fm.numElements++
}

func (hm *Hashmap) Add(key string, value int) {
	posic := hm.getBucketIndex(key)

	fmt.Printf("%d\n", posic)

	if hm.buckets[posic] == nil {
		hm.addWithNoCollision(key, value, posic)
	} else {
		hm.addWithCollision(key, value, posic)
	}
}

func (hm *Hashmap) Size() uint {
	return hm.numElements
}

func (hm *Hashmap) Get(key string) (int, bool) {
	posic := hm.getBucketIndex(key)
	existingBucket := hm.buckets[posic]

	// Empty bucket
	if existingBucket == nil {
		return 0, true
	}

	// Matches with the head
	if existingBucket.key == key {
		return existingBucket.value, false
	}

	// Look for the element in the linked list
	for existingBucket.next != nil {
		existingBucket = existingBucket.next
		if existingBucket.key == key {
			return existingBucket.value, false
		}
	}

	return 0, true
}

func (hm *Hashmap) GetKeys() []string {
	keys := make([]string, 0)

	for _, bucket := range hm.buckets {
		if bucket != nil {
			keys = append(keys, bucket.key)

			for bucket.next != nil {
				bucket = bucket.next
				keys = append(keys, bucket.key)
			}
		}
	}

	return keys
}
