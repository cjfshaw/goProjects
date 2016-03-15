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

func TestPutWithCollision(t *testing.T) {
	testElements := createElements()

	t.Log("Testing that put places the entry in the next empty space when a collision occurs.")

	testElements.Hashmap.put(1, 1)
	t.Logf("Hashmap after first put: %v", testElements.Hashmap)
	testElements.Hashmap.put(2, 1)
	t.Logf("Hashmap after second put: %v", testElements.Hashmap)
	testElements.Hashmap.put(3, 1)
	t.Logf("Hashmap after third put: %v", testElements.Hashmap)

	if testElements.Hashmap.Entries[2].Key != 3 {
		t.Errorf("Error: Put did not properly place the entry in a case where collisions occur.")
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

func TestGetWithCollission(t *testing.T) {
	testElements := createElements()

	t.Log("Testing get when the element it is searching for is not found in it's immediate index.")
	t.Log("Putting two elements into the map, who both want to be hashed at index 2.  It is expected that one will end up at index 2, and one will end up at index 0")

	testElements.Hashmap.put(5, 5)
	t.Logf("Hashmap after first put: %v", testElements.Hashmap)
	testElements.Hashmap.put(8, 8)
	t.Logf("Hashmap after second put: %v", testElements.Hashmap)

	value1, getErr1 := testElements.Hashmap.get(5)
	value2, getErr2 := testElements.Hashmap.get(8)
	index1 := testElements.Hashmap.hashToGetIndex(5)
	index2, matchErr := testElements.Hashmap.findIndexOfMatchingKey(8)

	if value1 != 5 {
		t.Errorf("Error: Value of first get (initial put) is not 5. \nValue: %v", value1)
	}

	if value2 != 8 {
		t.Errorf("Error: Value of second get (second put) is not 8. \nValue: %v", value2)
	}

	if index1 != 2 {
		t.Errorf("Error: Index of first put Entry is not 2.\nIndex: %v", index1)
	}

	if index2 != 0 {
		t.Errorf("Error: Index of second put Entry is not 0.\nIndex: %v", index2)
	}

	if getErr1 != nil {
		t.Errorf("Error: getErr1 is not nil.\n%v", getErr1)
	}

	if getErr2 != nil {
		t.Errorf("Error: getErr2 is not nil.\n%v", getErr2)
	}

	if matchErr != nil {
		t.Errorf("Error: matchErr is not nil.\n%v", matchErr)
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

func TestRemoveOnNil(t *testing.T) {
	testElements := createElements()

	t.Log("Testing remove on an empty map.")

	err := testElements.Hashmap.remove(8)

	if err == nil {
		t.Errorf("Error: err was nil.")
	}
}

func TestRemoveWithCollision(t *testing.T) {
	testElements := createElements()

	t.Log("Testing remove functionality when it is removing for a collision case.")

	testElements.Hashmap.put(5, 5)
	t.Logf("Hashmap after first put: %v", testElements.Hashmap)
	testElements.Hashmap.put(8, 8)
	t.Logf("Hashmap after second put: %v", testElements.Hashmap)
	err := testElements.Hashmap.remove(8)
	t.Logf("Hashmap after remove of key(8): %v", testElements.Hashmap)

	if testElements.Hashmap.NumEntries != 1 {
		t.Errorf("Error: NumEntries was not decremented. NumEntries: %v", testElements.Hashmap.NumEntries)
	}

	if testElements.Hashmap.Entries[0] != nil {
		t.Errorf("Error: Entry was not removed.")
	}

	if err != nil {
		t.Errorf("Error: err was not nil.\n%v", err)
	}
}

func TestHashToGetIndex(t *testing.T) {
	testElements := createElements()

	t.Log("Testing whether hashToGetIndex returns expected values.")

	length := len(testElements.Hashmap.Entries)

	t.Logf("Length of entries: %v", length)

	lessThan := testElements.Hashmap.hashToGetIndex(2)
	greaterThan := testElements.Hashmap.hashToGetIndex(5)
	equalTo := testElements.Hashmap.hashToGetIndex(3)

	if lessThan != 1 {
		t.Errorf("Error: 1 mod 3 should return 1 but it did not. Value: %v", lessThan)
	}

	if greaterThan != 2 {
		t.Errorf("Error: 5 mod 3 should return 2 but it did not. Value: %v", greaterThan)
	}

	if equalTo != 0 {
		t.Errorf("Error: x mod x should return 0 but it did not. Value: %v", equalTo)
	}
}

func TestFirstEmptySpaceLooping(t *testing.T) {
	testElements := createElements()

	t.Log("Testing that findFirstEmptySpace loops appropriately past the final element and returns the correct index.")

	testElements.Hashmap.put(3, 3)
	testElements.Hashmap.put(5, 3)

	hash3 := testElements.Hashmap.hashToGetIndex(3)
	hash5 := testElements.Hashmap.hashToGetIndex(5)

	t.Logf("Current hashmap %+v", testElements.Hashmap)
	t.Logf("Value of hash3: %v", hash3)
	t.Logf("Value of hash5: %v", hash5)

	firstEmptySpace, err := testElements.Hashmap.findFirstEmptySpace(2)

	if firstEmptySpace != 1 {
		t.Errorf("Error: The first empty space found was not at index 1.\nfirstEmptySpace: %v", firstEmptySpace)
	}

	if err != nil {
		t.Errorf("Error: The error was not nil, despite there being empty spaces in the map.\nerr: %v", err)
	}
}

func TestFirstEmptySpaceOnInitialIndex(t *testing.T) {
	testElements := createElements()

	t.Log("Testing that findFirstEmptySpace works when the first empty space is the same as the input.")

	testElements.Hashmap.put(3, 3)
	testElements.Hashmap.put(5, 3)

	hash3 := testElements.Hashmap.hashToGetIndex(3)
	hash5 := testElements.Hashmap.hashToGetIndex(5)

	t.Logf("Current hashmap %+v", testElements.Hashmap)
	t.Logf("Value of hash3: %v", hash3)
	t.Logf("Value of hash5: %v", hash5)

	firstEmptySpace, err := testElements.Hashmap.findFirstEmptySpace(1)

	if firstEmptySpace != 1 {
		t.Errorf("Error: The first empty space found was not at index 1.\nfirstEmptySpace: %v", firstEmptySpace)
	}

	if err != nil {
		t.Errorf("Error: The error was not nil, despite there being empty spaces in the map.\nerr: %v", err)
	}
}

func TestMatchNotFound(t *testing.T) {
	testElements := createElements()

	testElements.Hashmap.put(3, 3)
	testElements.Hashmap.put(5, 3)

	hash3 := testElements.Hashmap.hashToGetIndex(3)
	hash5 := testElements.Hashmap.hashToGetIndex(5)

	t.Logf("Current hashmap %+v", testElements.Hashmap)
	t.Logf("Value of hash3: %v", hash3)
	t.Logf("Value of hash5: %v", hash5)

	matchedIndex, err := testElements.Hashmap.findIndexOfMatchingKey(2)

	if matchedIndex != 0 {
		t.Errorf("Error: There should be no match found, so matched index should be 0. Matched Index: %v", matchedIndex)
	}

	if err == nil {
		t.Errorf("Error: err should not be nil, because there is no match.\n %v", err)
	}
}

func TestMatchFound(t *testing.T) {
	testElements := createElements()

	testElements.Hashmap.put(3, 3)
	testElements.Hashmap.put(5, 3)
	testElements.Hashmap.put(4, 3)

	hash3 := testElements.Hashmap.hashToGetIndex(3)
	hash4 := testElements.Hashmap.hashToGetIndex(4)
	hash5 := testElements.Hashmap.hashToGetIndex(5)

	t.Logf("Current hashmap %+v", testElements.Hashmap)
	t.Logf("Value of hash3: %v", hash3)
	t.Logf("Value of hash4: %v", hash4)
	t.Logf("Value of hash5: %v", hash5)

	matchedIndex, err := testElements.Hashmap.findIndexOfMatchingKey(4)

	if matchedIndex != 1 {
		t.Errorf("Error: The matched index was not 1. Matched Index: %v", matchedIndex)
	}

	if err != nil {
		t.Errorf("Error: err should be nil, because the match should occur on the index input.\n %v", err)
	}
}
