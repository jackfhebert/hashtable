

package HashTable

import (
  "List"
  )


type HashTable struct {
  size int
  capacity int
  // array of linkedlist pointers  
}

type slot struct {


  items *List.LinkedList
}