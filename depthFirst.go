package main

import (
  "fmt"
)

type Node struct {
  value int
  children []Node //every array is a pointer
}

func(n Node) DFS(){
  fmt.Println(n.value)
  for _,child:= range n.children{
    child.DFS()
  }
}

func main(){
  f := Node{9, []Node{} }
  e := Node{8, []Node{f} }
  d := Node{3, []Node{} }
  c := Node{12, []Node{} }
  b := Node{5, []Node{d,e} }
  a := Node{10, []Node{b,c} }
  a.DFS()
}
