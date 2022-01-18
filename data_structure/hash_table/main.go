package main


//const
const ArraySize = 7
// HashTable structure
type HashTable struct{
  Array [ArraySize]*bucket
}
// Insert
//TODO
func (h *HashTable) Insert(key string){
  index := hash(key)
  h.Array[index].Insert(key)
}
// Search
// TODO
func (h *HashTable) Search(key string) bool{
  index := hash(key)

}

// Delete
// TODO
func (h *HashTable) Delete(key string) bool{
  index := hash(key)

}

// bucket structure
type bucket struct{
  Head *bucketNode

}
// bucket Node structure is a linked list node that holds the key
type bucketNode struct{
  Key string
  next *bucketNode
}

// bucket.Insert takes in a key, creates a node with the key and inserts the node in the bucket
func (b *bucket) Insert(key string){
  newNode:= &bucketNode{Key: key}
  newNode.next = b.Head
  b.Head = newNode
}

// Search
//TODO
// Delete
//TODO

// hash
func hash(key string) int{
  var sum int
  for _,v:= range key{
    sum+=int(v)
  }
  return sum % ArraySize
}

// Init
func Init() *HashTable{
  result := &HashTable{}
  for i := range result.Array{
    result.Array[i] = &bucket{}
  }
  return result
}
func main(){

}
