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
	G.EdgeMap=make(map[int]Matrix,G.NumberOfVertex)
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
	fmt.Println(G.EdgeMap)
	fmt.Println(G.EdgeMap)
	lix:= G.PrimsMST()
	fmt.Println(G.EdgeList)
	fmt.Println(lix)



}

type Graph struct {
	NumberOfVertex int
	NumberOfEdge int
	EdgeList Matrix
	EdgeMap map[int]Matrix

}

func (g *Graph)AddEdge(li []int ){
	g.EdgeList=append(g.EdgeList,li)

	g.EdgeMap[li[0]]=append(g.EdgeMap[li[0]],[]int{li[1],li[2]})
	g.EdgeMap[li[1]]=append(g.EdgeMap[li[1]],[]int{li[0],li[2]})


}

func (g *Graph)PrimsMST()(Matrix){


	ll:=make([]bool,0)

	for i:=0; i<g.NumberOfEdge;i++{
		ll=append(ll,false)
	}

	sort.Sort(g.EdgeList)

	fmt.Println("sorted : ",g.EdgeList)
	ansList:=make(Matrix,0)

	li:= make(Matrix,0)

	li=append(li,g.EdgeList[0])

	ansList=append(ansList,li[0])
	fmt.Println("g.EdgeMap : ",g.EdgeMap)
	fmt.Println("li : ",li)
	fmt.Println("g.EdgeMap[li[0][0]] : ",g.EdgeMap[li[0][0]])
	//fmt.Println("g.EdgeMap[li[0][0]] : ",g.EdgeMap[li[0][0]])
	fmt.Println("g.EdgeMap[li[0][1]] : ",g.EdgeMap[li[0][1]])
	fmt.Println("")
	for i,v := range g.EdgeMap[li[0][0]] {
		fmt.Println("for loop 1: ",i,v)
		li = append(li, []int{li[0][0],v[0],v[1]})
	}
	for i,v := range g.EdgeMap[li[0][1]] {
		fmt.Println("for loop 2: ",i,v)
		li = append(li, []int{li[0][1],v[0],v[1]})
	}
	fmt.Println("li : ",li)
	li=append(li[:0],li[1:]...)
	fmt.Println("li : ",li)



	for i:=0;i<g.NumberOfVertex;i++ {
		if len(li)==0{
			fmt.Println("break li")
			break
		}
		fmt.Println("for loop 3: ",li)
		sort.Sort(li)
		if Isthere([]int{li[0][0],li[0][0]}) {
			ansList=append(ansList,li[0])
			fmt.Println("append ans : ",li[0])
		}
		for i,v := range g.EdgeMap[li[0][0]] {
			fmt.Println("for loop 10: ",i,v)
			if Isthere2([]int{li[0][0],v[0]}) {

				li = append(li, []int{li[0][0],v[0],v[1]})
			fmt.Println("append li : ",li[0][0],v[0],v[1])
			}
		}
		li=append(li[:0],li[1:]...)


	}

	return ansList



}

// func (g *Graph)FindParent(i int)bool{ }
type Matrix [][]int
func (m Matrix)Len()int{return len(m)}
func (m Matrix)Less(i,j int)bool{
	return m[i][2]<m[j][2]
}
func (m Matrix)Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}
