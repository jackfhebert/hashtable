package HashTable

import (
       "testing"
)

func Test_HashTableConstructor(t *testing.T) {
  NewHashTableSized(3)
  NewHashTable()
}

func Test_ContainsItems(t *testing.T) {
  table := NewHashTable()
  table.AddItem("key1", 7)
  table.AddItem("key2", 8)
  if table.size != 2 {
    t.Error("table has the wrong size")
  }
  if table.items[getIndex("key1", len(table.items))].Size() != 1 {
    t.Error("item isn't in the right spot")
  }

  if !table.ContainsKey("key1") {
    t.Error("didn't find key1")
  }
  if table.ContainsKey("foo") {
    t.Error("Shouldn't have found this key.")
  }
}

func Test_resizeTable(t *testing.T) {
  smallTable := NewHashTableSized(3)
  smallTable.AddItem("key1", "abc")
  smallTable.AddItem("key2", "abc")
  smallTable.AddItem("key3", "abc")
  smallTable.AddItem("key4", "abc")
  smallTable.AddItem("key5", "abc")
  for i := 0; i < 100; i += 1 {
    smallTable.AddItem("popularKey", i)
  }

  if !smallTable.ContainsKey("key1") {
    t.Error("didn't find key1")
  }
  if !smallTable.ContainsKey("key5") {
    t.Error("didn't find key5")
  }
  if !smallTable.ContainsKey("popularKey") {
    t.Error("didn't find popularKey")
  }
  if smallTable.size != 105 {
    t.Error("table has wrong size after resizing")
  }

  smallTable.RemoveKey("popularKey")
  if smallTable.size != 5 {
    t.Error("table has wrong size after resizing")
  }
}