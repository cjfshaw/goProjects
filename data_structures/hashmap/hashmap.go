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
	var index int

	if key > len(hmap.Entries) {
		index = key % len(hmap.Entries)
	} else if len(hmap.Entries) > key {
		index = len(hmap.Entries) % key
	} else {
		index = 0
	}

	return index
}

func (hmap *hashmap) findFirstEmptySpace(index int) (int, error) {
	var err error
	var newIndex, count int
	fmt.Println("in findFirstEmptySpace")

	for i := index; i != i+1; i = (i + 1) % len(hmap.Entries) {
		fmt.Println("i", i)
		if hmap.Entries[i] == nil {
			fmt.Println("returning in nil")
			return i, err
		}
		count++
		if count >= len(hmap.Entries) {
			err := fmt.Errorf("Error: No empty space found in this map to put.")
			return newIndex, err
		}
	}
	fmt.Println("returning at end")
	return newIndex, err
}

func (hmap *hashmap) put(key int, value int) error {
	newEntry := entry{Key: key, Value: value}

	index := hmap.hashToGetIndex(key)
	fmt.Println("index", index)

	newIndex, err := hmap.findFirstEmptySpace(index)
	fmt.Println("newindex", newIndex)
	if err != nil {
		fmt.Println("err found")
		return err
	}

	hmap.Entries[newIndex] = &newEntry
	hmap.NumEntries++

	return err
}

func (hmap *hashmap) findIndexOfMatchingKey(key int) (int, error) {
	var err error
	var newIndex, count int

	index := hmap.hashToGetIndex(key)

	for i := index; i != i+1; i = (i + 1) % len(hmap.Entries) {
		if hmap.Entries[i] == nil {
			err := fmt.Errorf("Error: No matching entries containing this key found.")
			return newIndex, err
		} else {
			if hmap.Entries[i].Key == key {
				return i, err
			}
			count++
			if count >= len(hmap.Entries) {
				err := fmt.Errorf("Error: No matching entries containing this key found.")
				return newIndex, err
			}

			if i == 0 {
				i++
			}
		}
	}

	return newIndex, err
}

func (hmap *hashmap) get(key int) (int, error) {
	var value int

	index, err := hmap.findIndexOfMatchingKey(key)

	if hmap.Entries[index] == nil {
		err = fmt.Errorf("Error, value of desired key is nil.")
		return value, err
	}

	value = hmap.Entries[index].Value

	return value, err
}

func (hmap *hashmap) remove(key int) error { //find matching key in get and remove
	index, err := hmap.findIndexOfMatchingKey(key)

	if err != nil {
		return err
	}

	hmap.Entries[index] = nil
	hmap.NumEntries--

	return err
}

func main() {
	fmt.Println("main print")
}
