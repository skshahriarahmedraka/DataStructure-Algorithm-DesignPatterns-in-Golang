package main

import (
  "fmt"

)

type List struct{
  root Element
  len int
}

type Element struct {

  list *List
  next, previous *Element
  Value interface{}
}

func New() *List {
   return new(List).Init()
}

func (l *List) Init() *List {
  l.root.next= &l.root
  l.root.previous= &l.root
  l.len=0
}

func (l *List) Len() int {
  return l.len
}






func main() {


  }
