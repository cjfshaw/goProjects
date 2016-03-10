package hashmap

import "testing"

type testElements struct {
	Hashmap hashmap
	Entries []*entry
}

func createElements() testElements {
	entries := make([]*entry, 0)

	entry1 := entry{Key: 1, Value: 1}
	entry2 := entry{Key: 2, Value: 4}
	entry3 := entry{Key: 3, Value: 5}

	entries = append(entries, &entry1)
	entries = append(entries, &entry2)
	entries = append(entries, &entry3)

	mapEntries := make([]*entry, 3)
	newHashmap := hashmap{Entries: mapEntries}

	newElements := testElements{Hashmap: newHashmap, Entries: entries}

	return newElements
}

func TestPut(t *testing.T) {
	testElements := createElements()

	t.Log("Testing put into an empty list.")

	testElements.Hashmap.put(1, 1)

	if testElements.Hashmap.Entries[0].Key != 1 && testElements.Hashmap.Entries[0].Value != 1 {
		t.Errorf("Error: Entry is not in expected location.\n 0 Entry: %v", testElements.Hashmap.Entries[0])
	}
}

func TestGet(t *testing.T) {
	testElements := createElements()

	t.Log("Testing get on a map with one value.  No collisions.")

	testElements.Hashmap.put(1, 1)
	value := testElements.Hashmap.get(1)

	if value != 1 {
		t.Errorf("Error: Value found is not value put.  Value: %v", value)
	}
}
