/*
Simple HashTable implementation.
It auto-resizes and uses silly linked lists to store the items
per bucket. Built only to be more familar with golang.
*/

package HashTable

import (
	"fmt"
	"hash/adler32"
)

var (
	// When a table reaches this ratio of elements to number of
	// buckets, it will resize and double the number of buckets.  
	fillRate int = 10
)

// The externally visible implementation.
type HashTable struct {
	// How many items are current in the table.
	size int
	// The maximum number of items that the table should hold.
	// If calling AddItem increases 'size' past this value, then
	// the table will resize itself in place.
	capacity int
	// Array of linkedlist pointers.
	items []*LinkedList
}

// Internal helper to wrap the key and value which were added to
// the hashtable. This is the value stored in the linked lists
// per bucket.
type tableItem struct {
	key   string
	value interface{}
}

// Create a new hashtable with a given number of buckets. This
// probably should have been made to specify the capacity instead.
func NewHashTableSized(size int) *HashTable {
	table := &HashTable{0, fillRate * size, make([]*LinkedList, size)}
	for i := 0; i < len(table.items); i++ {
		table.items[i] = nil
	}
	return table
}

// Return a default hashtable.
func NewHashTable() *HashTable {
	return NewHashTableSized(128)
}

// Make the hashtable bigger so that there are fewer items per bucket.
// This takes up more memory (ish) but reduces the number of items per
// bucket which makes the datastructure faster to use.
func (table *HashTable) resizeTable() {
	next := NewHashTableSized(2 * len(table.items))
	for _, list := range table.items {
		if list != nil {
			for _, item := range list.Items() {
				if parsed, ok := (*item).(tableItem); ok {
					next.AddItem(parsed.key, parsed.value)
				} else {
					fmt.Println("failed to parse item in resize", item)
				}
			}
		}
	}
	table = next
}

// Helper function to take the key string and determine which bucket
// the item should be placed in. This is the simplest hash function
// I found in the core libraries. Not really sure how appropriate it is.
func getIndex(key string, max int) int {
	hash := adler32.New()
	hash.Write([]byte(key))
	digest := hash.Sum32()

	return int(digest) % max
}

// Add a key, value paid to the hash table.
func (table *HashTable) AddItem(key string, value interface{}) {
	index := getIndex(key, len(table.items))
	if table.items[index] == nil {
		table.items[index] = NewLinkedList()
	}
	table.size += 1
	table.items[index].AddItem(tableItem{key, value})

	if table.size > table.capacity {
		table.resizeTable()
	}
}

// Remove all instances of a key from the table.
func (table *HashTable) RemoveKey(key string) bool {
	index := getIndex(key, len(table.items))
	if table.items[index] != nil {
		for _, item := range table.items[index].Items() {
			if parsed, ok := (*item).(tableItem); ok {
				if parsed.key == key {
					table.items[index].RemoveItem(item)
					table.size -= 1
				}
			}
		}
	}
	return false
}

// Determine if a key is contained in the hash table.
func (table *HashTable) ContainsKey(key string) bool {
	fmt.Println("Checking for key:", key)
	index := getIndex(key, len(table.items))

	if table.items[index] != nil {
		for _, item := range table.items[index].Items() {
			if parsed, ok := (*item).(tableItem); ok {
				if parsed.key == key {
					fmt.Println(parsed.key)
					return true
				}
			} else {
				fmt.Println("failed to parse item in contains", *item)
			}
		}
	}
	return false
}

func (table *HashTable) GetValue(key string) interface{} {
	fmt.Println("Checking for key:", key)
	index := getIndex(key, len(table.items))

	if table.items[index] != nil {
		for _, item := range table.items[index].Items() {
			if parsed, ok := (*item).(tableItem); ok {
				if parsed.key == key {
					return parsed.value
				}
			} else {
				fmt.Println("failed to parse item in contains", *item)
			}
		}
	}
	return nil
}