/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */

package main

import "fmt"


type Doubly struct {
	Head *Node
}

func NewDoubly() *Doubly {
	return &Doubly{nil}
}

// AddAtBeg Add a node to the beginning of the linkedlist
func (ll *Doubly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head

	if ll.Head != nil {
		ll.Head.Prev = n
	}

	ll.Head = n

}

// AddAtEnd Add a node at the end of the linkedlist
func (ll *Doubly) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	n.Prev = cur
}

// DelAtBeg Delete the node at the beginning of the linkedlist
func (ll *Doubly) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next

	if ll.Head != nil {
		ll.Head.Prev = nil
	}
	return cur.Val
}

// DetAtEnd Delete a node at the end of the linkedlist
func (ll *Doubly) DelAtEnd() interface{} {
	// no item
	if ll.Head == nil {
		return -1
	}

	// only one item
	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	// more than one, go to second last
	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	return retval
}

// Count Number of nodes in the linkedlist
func (ll *Doubly) Count() interface{} {
	var ctr int = 0

	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}

	return ctr
}

// Reverse Reverse the order of the linkedlist
func (ll *Doubly) Reverse() {
	var Prev, Next *Node
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
	}

	ll.Head = Prev
}

// Display diplay the linked list
func (ll *Doubly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}

// DisplayReverse Display the linkedlist in reverse order
func (ll *Doubly) DisplayReverse() {
	if ll.Head == nil {
		return
	}
	var cur *Node
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {
	}

	for ; cur != nil; cur = cur.Prev {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}

