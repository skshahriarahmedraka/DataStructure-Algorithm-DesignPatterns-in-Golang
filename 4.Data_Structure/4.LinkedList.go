/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */

package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}


func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

type  Singly struct {
	length  int
	Head *Node
}


func NewSingly() *Singly{
	return &Singly{0,nil}
}


func (l *Singly) AddAtFront(val interface{}){
	n:= NewNode(val)
	n.Next =l.Head
	l.Head =n
	l.length++
}

func (l *Singly) AddAtEnd(val int ){
	n:= NewNode(val)
	if l.Head == nil {
		l.Head = n
		l.length++
		return
	}
	cur := l.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	l.length++
}

func (l *Singly) DelAtBeg() interface{} {
	if l.Head == nil {
		return -1
	}
	cur := l.Head
	l.Head = cur.Next
	l.length--
	return cur.Val
}

func (l *Singly) DelAtEnd() interface{} {
	if l.Head == nil {
		return -1
	}

	if l.Head.Next == nil {
		return l.DelAtBeg()
	}

	cur := l.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	l.length--
	return retval

}
func (l *Singly) Reverse() {
	var prev, Next *Node
	cur := l.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}

	l.Head = prev
}
func (ll *Singly) ReversePartition(left, right int) error {
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		return err
	}
	tmpNode := NewNode(-1)
	tmpNode.Next = ll.Head
	pre := tmpNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	ll.Head = tmpNode.Next
	return nil
}
func (ll *Singly) CheckRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}
	return nil
}
func (ll *Singly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}


func main (){
	fmt.Println("********singly linked list *************")
	l:= NewSingly()
	l.AddAtEnd(12)
	l.AddAtEnd(23)
	l.AddAtEnd(99)

	l.Display()

}


