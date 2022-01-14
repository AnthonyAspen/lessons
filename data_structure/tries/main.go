package main

import "fmt"

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

// initTrie creates a new trie 
func InitTrie() (result *Trie){
  result = &Trie{root: &Node{}}
  return 
}

// Insert takes in a word and adds it to the trie
func (t *Trie) Insert(w string){
  wordLength := len(w)
  currentNode := t.root
  for i := 0;i<wordLength;i++{
    charIndex := w[i] - 'a'
    if currentNode.childern[charIndex] == nil{
      currentNode.childern[charIndex] = &Node{}
    }
    currentNode = currentNode.childern[charIndex]
  }
  currentNode.isEnd = true
}

// Search takes in a word and RETURNs true if the word is included in the trie
func (t Trie) Search(w string) bool{
  wordLength := len(w)
  currentNode := t.root
  for i := 0;i < wordLength;i++{
    charIndex := w[i] - 'a'
    if currentNode.childern[charIndex] == nil{
      return false
    }
    currentNode = currentNode.childern[charIndex]
  }
  if currentNode.isEnd == true{
    return true
  }
  // if currentNode.isEnd != true
  return false
}

func main(){
  trie := InitTrie()
  trie.Insert("orc")
  trie.Insert("orb")
  fmt.Println(trie.Search("orc"))
  fmt.Println(trie.Search("orb"))
  fmt.Println(trie.Search("ora"))

}
