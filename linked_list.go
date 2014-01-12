package List


type LinkedListNode struct {
  // The value of this node.
  value string
  // Pointer to the next node in the list.
  next *LinkedListNode
}

type LinkedList struct {
  // The first node in the list.
  first *LinkedListNode
  // The last item in the list.
  last *LinkedListNode
  // How many items are in the list.
  size int
}

func NewLinkedList() *LinkedList {
  return &LinkedList{nil, nil, 0}
}

func (list *LinkedList) Size() int {
  return list.size
}

func (list *LinkedList) AddItem(item string) {
  list.size += 1
  node := &LinkedListNode{item, nil}
  if list.first == nil {
    list.first = node;
  }
  if list.last != nil {
    list.last.next = node
  }
  list.last = node  
}

func (list *LinkedList) RemoveItem(item string) {
  var prevNode *LinkedListNode
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
  
    prevNode = currNode  
  }
}