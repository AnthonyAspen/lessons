package main


//const
const ArraySize = 7
// HashTable structure
type HashTable struct{
  array [ArraySize]*bucket
}
// Insert
// Search
// Delete

// bucket structure
type bucket struct{
  head *bucketNode

}
// bucket Node structure
type bucketNode struct{

}

// Insert
// Search
// Delete

// hash

// Init
func Init() *HashTable{
  result := &HashTable{}
  for i := range result.array{
    result.array[i] = &bucket{}
  }
  return result
}
func main(){

}
