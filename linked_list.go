/*
Package List provides a basic linked list.
Obviously this isn't strictly needed, but you have to learn somehow.

Adding items - O(c)
Removing items - O(N)
Checking size - O(c)
*/

package HashTable


// Not publicly visible since it is our internal wrapper.
type linkedListNode struct {
  // The value of this node.
  // Once I figure out interfaces, that will go here.
  value interface{}
  // Pointer to the next node in the list.
  next *linkedListNode
}

// Exposed - this is the struct to use.
type LinkedList struct {
  // The first node in the list.
  first *linkedListNode
  // The last item in the list.
  // This makes adding to the list fast, but isn't strictly
  // needed.
  last *linkedListNode
  // How many items are in the list. This is mostly
  // for the size helper and not strictly needed.
  size int
}

func NewLinkedList() *LinkedList {
  return &LinkedList{nil, nil, 0}
}

func (list *LinkedList) Size() int {
  return list.size
}

func (list *LinkedList) AddItem(item interface{}) {
  list.size += 1
  node := &linkedListNode{item, nil}
  if list.first == nil {
    list.first = node;
  }
  if list.last != nil {
    list.last.next = node
  }
  list.last = node  
}

func (list *LinkedList) RemoveItem(item interface{}) {
  // Track the previous node from the iterator for updating
  // pointers between nodes.
  var prevNode *linkedListNode
  prevNode = nil
  for currNode := list.first; currNode != nil; currNode = currNode.next {
    if currNode.value == item {
      // Update the list metadata.
      list.size -= 1
      if currNode == list.first {
        list.first = currNode.next
      }
      if currNode == list.last {
        list.last = prevNode
      }
      // Update the nodes.
      if prevNode != nil {
        prevNode.next = currNode.next
      }

      // All done here.
      return
    }

    // Keep iterating through the list. I could probably assign
    // this in the for-loop definition above.  
    prevNode = currNode  
  }
}

func (list *LinkedList) Items() []*interface{} {
  items := make([]*interface{}, list.size)
  
  for i, currNode := 0, list.first; currNode != nil; currNode = currNode.next {
    items[i] = &currNode.value
    i += 1
  }
  return items
}