package hashmap

import (
	"fmt"
	"hash/fnv"
)

type bucket struct {
	key   string
	value int
	next  *bucket
}

type Hashmap struct {
	buckets     []bucket
	numElements int
}

type Hashmaper interface {
	Add(key string, value int)
	Remove(key string) bool
	Get(key string) int
	GetKeys() []string
}

func NewHashmap(numBuckets int) *Hashmap {
	hashmap := new(Hashmap)
	hashmap.buckets = make([]bucket, numBuckets)
	hashmap.numElements = 0
	return hashmap
}

func getHashCode(key string) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(key))
	return algorithm.Sum32()
}

func (fm *Hashmap) getBucketIndex(key string) uint32 {
	hashCode := getHashCode(key)
	idx := hashCode % uint32(cap(fm.buckets))
	return idx
}

func (hm *Hashmap) Add(key string, value int) {
	// idx := hm.getBucketIndex(key)
	fmt.Printf("%#v", hm.buckets)

	// int posic = getBucketIndex(key);
	// if (this.bucket[posic] == null) {   // No collision
	// 	addNoCollision(key, value, posic);
	// } else { // Collision
	// 	addCollision(key, value, posic);
	// }
}
