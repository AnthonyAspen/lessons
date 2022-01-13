package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct{
  Key int
  Left *Node
  Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert (k int )(error){
  // if k is larger we should move right, otherwise move left
  if n.Key < k{
    //move right
    // if the right node is empty, the key takes the place, otherwise go deeper
    if n.Right == nil{
      n.Right = &Node{Key:k}
    } else {
      n.Right.Insert(k)

    }
  }else if n.Key>k{
    //move left
    if n.Left == nil{
      n.Left = &Node{Key:k}
    } else{
      n.Left.Insert(k)
    }
  }

  return nil
}


// Search
// returns true if there is a node with that value
func (n *Node) Search(k int) bool{
  // the the node is empty then return false, didn't find the match
  if n == nil{
    return false
  }
  if n.Key < k{
    return n.Right.Search(k)
  }else if n.Key > k{
    return n.Left.Search(k)
  }
  return true
}

//Display
func Display(n *Node){
  if n == nil {
    return 
  }
  Display(n.Left)
  fmt.Print(n.Key," ")
  Display(n.Right)
}

// Main func
func main(){
  tree := &Node{Key:100}
  for i := 0;i<10;i++{
    tree.Insert(i*i+i)
  }
  Display(tree)


}



