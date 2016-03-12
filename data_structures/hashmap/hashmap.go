package hashmap

import "fmt"

type entry struct {
	Key   int
	Value int
}

type hashmap struct {
	Entries    []*entry
	NumEntries int
}

func (hmap *hashmap) hashToGetIndex(key int) int {
	index := len(hmap.Entries) % key

	return index
}

func (hmap *hashmap) put(key int, value int) {
	newEntry := entry{Key: key, Value: value}

	index := hmap.hashToGetIndex(key)

	hmap.Entries[index] = &newEntry
	hmap.NumEntries++
}

func (hmap *hashmap) get(key int) (int, error) {
	var err error
	var value int
	index := hmap.hashToGetIndex(key)

	if hmap.Entries[index] == nil {
		err = fmt.Errorf("Error, value of desired key is nil.")
		return value, err
	}

	value = hmap.Entries[index].Value

	return value, err
}

func (hmap *hashmap) remove(key int) {
	index := hmap.hashToGetIndex(key)

	hmap.Entries[index] = nil
	hmap.NumEntries--
}

func main() {
	fmt.Println("main print")
}
