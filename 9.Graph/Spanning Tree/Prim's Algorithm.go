package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main (){
	var t,t2 int
	fmt.Printf("number of Vertices : ")
	fmt.Scanln(&t)
	fmt.Printf("number of Edge : ")
	fmt.Scanln(&t2)
	G:= Graph{NumberOfVertex: t,NumberOfEdge: t2}
	for i :=0 ;i<t2;i++ {
		var li []int
		fmt.Printf("give - Vertex1 , Verex2 , Weight : ")
		b,_,_:= bufio.NewReader(os.Stdin).ReadLine()
		sli:= strings.Split(string(b)," ")
		for j:=0 ;j<len(sli);j++{
			temp,_:=strconv.Atoi(sli[j])
			li = append(li, temp)
		}
		G.AddEdge(li)


	}
	fmt.Println(G.EdgeList)
	lix:= G.PrimsMST()
	fmt.Println(G.EdgeList)
	fmt.Println(lix)



}

type Graph struct {
	NumberOfVertex int
	NumberOfEdge int
	EdgeList Matrix
}

func (g *Graph)AddEdge(li []int ){
	g.EdgeList=append(g.EdgeList,li)
}

func (g *Graph)PrimsMST()([][]int){


}

func (g *Graph)FindParent(i int)bool{



}
type Matrix [][]int
func (m Matrix)Len()int{return len(m)}
func (m Matrix)Less(i,j int)bool{
	return m[i][2]<m[j][2]
}
func (m Matrix)Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}