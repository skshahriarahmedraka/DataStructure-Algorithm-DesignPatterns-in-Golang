/**
 * @Author: sk shahriar ahmed raka <ssar>
 * @Date:   01-Jan-1970
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: binarysearchTree.go
 * @Last modified by:   ssar
 * @Last modified time: 27-Mar-2021
 */



package main


import (
  "fmt"
)

type Node struct {


}

type BinaryTree struct{
  Data string
  LeftNode *BinaryTree
  RightNode *BinaryTree  // Node
  ParentNode *BinaryTree
}

func (B *BinaryTree) Insert(data string ){
  if B.Data == "" && B.LeftNode==nil && B.RightNode==nil {
    B.Data=data
    // B = &Node(Data:data , LeftNode : nil ,RightNode : nil,ParentNode:nil)
  }else {
    run :=true
    extranode = &B
    for (run) {
      if data <= extranode.Data {
        if extranode.LeftNode == nil {
          // B = &Node(Data:data , LeftNode : nil ,RightNode : nil,ParentNode:nil)
            extranode.LeftNode = &BinaryTree{Data: data , LeftNode : nil ,RightNode : nil,ParentNode: extranode}
          extranode = extranode.LeftNode

        }
      }else if data > extranode.Data {
        if extranode.RightNode == nil {
          // B = &Node(Data:data , LeftNode : nil ,RightNode : nil,ParentNode:nil)
            extranode.RightNode = &BinaryTree{Data:data , LeftNode : nil ,RightNode : nil,ParentNode: extranode}
            fmt.Println("inserted data in RightNode")
            run =false
        }else {
          extranode = extranode.RightNode

        }
      }
    }
  }
}

func (B *BinaryTree) ValueList() {
    if B.Value==nil {
      fmt.Println("")
      return
    }

  }




func main(){

}
