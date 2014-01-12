package List

import (
       "testing"
       )

func Test_Constructor(t *testing.T) {
  NewLinkedList()
  t.Log("success.")
}

func Test_AddItems(t* testing.T) {
  list := NewLinkedList()
  if list.Size() != 0 {
    t.Error("Incorrect size.")
  }
  list.AddItem("foo")
  list.AddItem("bar")
  list.AddItem("baz")
  if list.Size() != 3 {
    t.Error("Incorrect size.")
  }

  list.RemoveItem("cat")
  if list.Size() != 3 {
    t.Error("Incorrect size.")
  }
  list.RemoveItem("bar")
  list.RemoveItem("foo")
  list.RemoveItem("baz")
  if list.Size() != 0 {
    t.Error("Incorrect size.")
  }
}