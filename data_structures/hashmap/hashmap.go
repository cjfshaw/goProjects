package hashmap

import "fmt"

type entry struct {
	Key   int
	Value int
}

type hashmap struct {
	Entries []*entry
}

func (hmap *hashmap) hashToGetIndex(key int) int {
	index := len(hmap.Entries) % key

	return index
}

func (hmap *hashmap) put(key int, value int) {
	newEntry := entry{Key: key, Value: value}

	index := hmap.hashToGetIndex(key)

	hmap.Entries[index] = &newEntry
}

func (hmap *hashmap) get(key int) int {
	index := hmap.hashToGetIndex(key)

	value := hmap.Entries[index].Value

	return value
}

func main() {
	fmt.Println("main print")
}
