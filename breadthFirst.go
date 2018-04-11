package main

import (
  "fmt"
)

type Node struct {
  value int
  children []Node
}

func BFS(queue []Node){
  n := queue[0]
  fmt.Println(n.value)
  for _,child:=range n.children{
    queue = append(queue, child)
  }
  if len(queue) != 1{
    BFS(queue[1:])
  }
}


func main(){
  f := Node{9, []Node{} }
  e := Node{8, []Node{f} }
  d := Node{3, []Node{} }
  c := Node{12, []Node{} }
  b := Node{5, []Node{d,e} }
  a := Node{10, []Node{b,c} }
  BFS([]Node{a})
}
