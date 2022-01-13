package main


// AlphabetSize is the number of possible characters in the trrie
const AlphabetSize = 26
// Node
type Node struct{
  childern [AlphabetSize]*Node
  isEnd bool
}

// Trie
type Trie struct{
  root *Node
}

func InitTrie() *Trie{

  return 
}

// Insert

// Search

func main(){

}
