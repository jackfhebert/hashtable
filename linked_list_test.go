package HashTable

import (
	"testing"
)

// Basic test. Seems to get called.
func Test_Constructor(t *testing.T) {
	NewLinkedList()
	t.Log("success.")
}

// Test some basic list functionality.
func Test_AddItems(t *testing.T) {
	// List starts off ok.
	list := NewLinkedList()
	if list.Size() != 0 {
		t.Error("Incorrect size.")
	}

	// We add items and the size increases as expected.
	list.AddItem("foo")
	list.AddItem("bar")
	list.AddItem("baz")
	if list.Size() != 3 {
		t.Error("Incorrect size.")
	}

	// Try removing an item not in the list, no errors.
	list.RemoveItem("cat")
	if list.Size() != 3 {
		t.Error("Incorrect size.")
	}
	// Remove items in a random order.
	list.RemoveItem("bar")
	list.RemoveItem("foo")
	list.RemoveItem("baz")
	// List ends up at the correct size.
	if list.Size() != 0 {
		t.Error("Incorrect size.")
	}
}

// Test how duplicates are handled.
func Test_Duplicates(t *testing.T) {
	list := NewLinkedList()
	list.AddItem("first")
	list.AddItem("second")
	list.AddItem("second")
	if list.Size() != 3 {
		t.Error("Incorrect size.")
	}

	list.RemoveItem("second")
	if list.Size() != 2 {
		t.Error("Incorrect size.")
	}
	list.RemoveItem("second")
	if list.Size() != 1 {
		t.Error("Incorrect size.")
	}
}

func Test_MultipleTypes(t *testing.T) {
	list := NewLinkedList()
	list.AddItem("foo")
	list.AddItem(17)
	if list.Size() != 2 {
		t.Error("Incorrect size.")
	}

	items := list.Items()
	if len(items) != 2 {
		t.Error("Incorrect size.")
	}
	if *items[0] != "foo" {
		t.Error("Incorrect items.")
	}
	if *items[1] != 17 {
		t.Error("Incorrect items.")
	}

	list.AddItem(list)
	if list.Size() != 3 {
		t.Error("Incorrect size.")
	}
}
