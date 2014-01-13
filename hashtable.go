

package HashTable


import (
  "fmt"
  "hash/adler32"
)

var (
  fillRate int = 10
)

type HashTable struct {
  // How many items are current in the table.
  size int
  // The maximum number of items that the table should hold.
  capacity int
  // array of linkedlist pointers  
  items []*LinkedList
}

type tableItem struct {
  key string
  value interface{}
}

func NewHashTableSized(size int) *HashTable {
  table := &HashTable{0, fillRate * size, make([]*LinkedList, size)}
  for i := 0; i < len(table.items); i++ {
    table.items[i] = nil
  }
  return table
}

func NewHashTable() *HashTable {
  return NewHashTableSized(128)
}

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

func getIndex(key string, max int) int {
  hash := adler32.New()
  hash.Write([]byte(key))
  digest := hash.Sum32()

  return int(digest) % max
}

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

func (table* HashTable) RemoveKey(key string) bool {
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
