package main

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	set := &HashSet{
		Map: make(map[interface{}]bool),
	}

	// Testing Put
	set.Put("key1")
	if !set.Contains("key1") {
		t.Errorf("key1 should exist in the set")
	}

	// Re-Put the same key
	set.Put("key1")
	keys := set.Keys()
	if len(keys) != 1 {
		t.Errorf("set should only contain key1 once")
	}

	// Testing Contains
	if set.Contains("key2") {
		t.Errorf("key2 should not exist in the set")
	}

	// Testing Remove
	set.Put("key2")
	set.Remove("key2")
	if set.Contains("key2") {
		t.Errorf("key2 should have been removed from the set")
	}

	// Testing PutIfAbsent
	old, absent := set.PutIfAbsent("key1")
	if old != "key1" || absent != false {
		t.Errorf("key1 is already present in the set, so it should not be marked as absent")
	}

	old, absent = set.PutIfAbsent("key3")
	if old != "new string added!" || absent != true {
		t.Errorf("key3 was not in the set, so it should be marked as absent and added")
	}
	if !set.Contains("key3") {
		t.Errorf("key3 should have been added to the set")
	}

	// Testing Keys
	keys = set.Keys()
	expectedKeys := []string{"key1", "key3"}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Errorf("expected %s but got %s", expectedKeys[i], key)
		}
	}
}

// func main() {
// 	// Running the tests
// 	t := testing.T{}
// 	TestHashSet(&t)
// }
