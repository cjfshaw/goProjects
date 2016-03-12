package hashmap

import "testing"

type testElements struct {
	Hashmap hashmap
	Entries []*entry
	Keys    []int
	Values  []int
}

func createElements() testElements {
	entries := make([]*entry, 0)

	entry1 := entry{Key: 1, Value: 1}
	entry2 := entry{Key: 2, Value: 4}
	entry3 := entry{Key: 3, Value: 5}

	entries = append(entries, &entry1)
	entries = append(entries, &entry2)
	entries = append(entries, &entry3)

	keys := make([]int, 3)
	values := make([]int, 3)
	keys[0] = 1
	keys[1] = 2
	keys[2] = 3
	values[0] = 1
	values[1] = 4
	values[2] = 5

	mapEntries := make([]*entry, 3)
	newHashmap := hashmap{Entries: mapEntries, NumEntries: 0}

	newElements := testElements{Hashmap: newHashmap, Entries: entries, Keys: keys, Values: values}

	return newElements
}

func TestPut(t *testing.T) {
	testElements := createElements()

	t.Log("Testing put into an empty list.")
	originalNumEntries := testElements.Hashmap.NumEntries
	testElements.Hashmap.put(1, 1)
	newNumEntries := testElements.Hashmap.NumEntries

	if testElements.Hashmap.Entries[0].Key != 1 && testElements.Hashmap.Entries[0].Value != 1 {
		t.Errorf("Error: Entry is not in expected location.\n 0 Entry: %v", testElements.Hashmap.Entries[0])
	}

	if newNumEntries != (originalNumEntries + 1) {
		t.Errorf("Error: NumEntries did not increment by 1.\n Original number of entries was %v\n Current number of entries is %v", originalNumEntries, newNumEntries)
	}
}

func TestGet(t *testing.T) {
	testElements := createElements()

	t.Log("Testing get on a map with one value.  No collisions.")

	testElements.Hashmap.put(1, 1)
	value, err := testElements.Hashmap.get(1)

	if value != 1 {
		t.Errorf("Error: Value found is not value put.  Value: %v", value)
	}

	if err != nil {
		t.Errorf("Error: err is not nil when doing get on a valid key.\n %v", err)
	}
}

func TestGetOnNil(t *testing.T) {
	testElements := createElements()

	t.Log("Testing get on an empty map.  Expecting an error returned.")

	_, err := testElements.Hashmap.get(1)

	if err == nil {
		t.Errorf("Error: No error produced when trying to get on an empty map.")
	}

}

func TestRemove(t *testing.T) {
	testElements := createElements()

	t.Log("Testing remove functionality, expecting remove to work without error.")

	testElements.Hashmap.put(testElements.Keys[0], testElements.Values[0])
	testElements.Hashmap.remove(testElements.Keys[0])
	index := testElements.Hashmap.hashToGetIndex(testElements.Keys[0])
	value := testElements.Hashmap.Entries[index]

	t.Logf("Index is %v", index)

	if testElements.Hashmap.Entries[index] != nil {
		t.Errorf("Error: Value of hashmap at removed key's index is not nil.  Value: %v", value)
	}
}
