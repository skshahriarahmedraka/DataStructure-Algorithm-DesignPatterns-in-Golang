/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */

package main

import "fmt"

type Node struct {
	Value int
	LeftNode *Node
	RightNode *Node
	//ParentNode *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value,LeftNode: nil,RightNode: nil,ParentNode: nil}
}

type BinaryTree struct{
	Root *Node
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func CalculateDepth(n *Node , depth int)int {
	if n== nil {
		return  depth
	}
	return Max(CalculateDepth(n.LeftNode,depth+1),CalculateDepth(n.RightNode,depth+1))
}

func Insert(root *Node,val int )*Node{
	if root ==nil {
		return NewNode(val)
	}
	if val <root.Value{
		root.LeftNode=Insert(root.LeftNode,val)
	}else{
		root.RightNode=Insert(root.RightNode,val)
	}
	return root
}

func (t *BinaryTree)Depth()int {
	return CalculateDepth(t.Root ,0)
}

func InOrder(n *Node){
	if n!= nil {
		InOrder(n.LeftNode)
		fmt.Print(n.Value ," ")
		InOrder(n.RightNode)
	}
}

func BstDelete(root *Node,val int)*Node {
	if root ==nil {
		return nil
	}
	if val <root.Value {
		root.LeftNode = BstDelete(root.LeftNode,val)
	}else if val >root.Value{
		root.RightNode =BstDelete(root.RightNode,val)
	}else{
		if root.LeftNode==nil {
			return root.RightNode
		}else if root.RightNode ==nil {
			return root.LeftNode
		}else {
			n:= root.RightNode
			d:= InOrderSuccessor(n)
			d.
		}
	}
}


func main(){

}