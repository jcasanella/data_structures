package hashmap

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

func main() {
	var hashMap *Hashmap = NewHashmap(5)
	_ = hashMap
}
