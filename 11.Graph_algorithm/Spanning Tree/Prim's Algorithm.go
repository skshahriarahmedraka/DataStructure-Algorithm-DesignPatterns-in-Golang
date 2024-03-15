//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func main (){
//	var t,t2 int
//	fmt.Printf("number of Vertices : ")
//	fmt.Scanln(&t)
//	fmt.Printf("number of Edge : ")
//	fmt.Scanln(&t2)
//	G:= Graph{NumberOfVertex: t,NumberOfEdge: t2}
//	for i :=0 ;i<t2;i++ {
//		var li []int
//		fmt.Printf("give - Vertex1 , Verex2 , Weight : ")
//		b,_,_:= bufio.NewReader(os.Stdin).ReadLine()
//		sli:= strings.Split(string(b)," ")
//		for j:=0 ;j<len(sli);j++{
//			temp,_:=strconv.Atoi(sli[j])
//			li = append(li, temp)
//
//		}
//		if _,ok := G.EdgeMap[li[0]] ; ok  {
//			ll:=[]int{li[1],li[2]}
//			G.EdgeMap[li[0]]= append(G.EdgeMap[li[0]],ll)
//		}//else {
//		//	ll:=Matrix{[li[1],li[2]]}
//		//	G.EdgeMap[li[0]]=ll
//		//}
//		if _,ok := G.EdgeMap[li[1]] ; ok  {
//			ll:=[]int{li[0],li[2]}
//			G.EdgeMap[li[1]]= append(G.EdgeMap[li[1]],ll)
//		}//else {
//		//	ll:=Matrix{[li[1],li[2]]}
//		//	G.EdgeMap[li[0]]=ll
//		//}
//		G.AddEdge(li)
//
//
//	}
//	fmt.Println(G.EdgeList)
//	fmt.Println(G.EdgeMap)
//	lix:= G.PrimsMST()
//	fmt.Println(G.EdgeList)
//	fmt.Println(lix)
//
//
//
//}
//
//type Graph struct {
//	NumberOfVertex int
//	NumberOfEdge int
//	EdgeList Matrix
//	EdgeMap map[int]Matrix
//
//}
//
//func (g *Graph)AddEdge(li []int ){
//	g.EdgeList=append(g.EdgeList,li)
//}
//
//func (g *Graph)PrimsMST()([][]int){
//
//
//}
//
//func (g *Graph)FindParent(i int)bool{
//
//
//
//}
//type Matrix [][]int
//func (m Matrix)Len()int{return len(m)}
//func (m Matrix)Less(i,j int)bool{
//	return m[i][2]<m[j][2]
//}
//func (m Matrix)Swap(i,j int){
//	m[i],m[j]=m[j],m[i]
//}
//

//////////////////////////////////////////////////////////////////////////////////////

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func main (){
//	var t,t2 int
//	fmt.Printf("number of Vertices : ")
//	fmt.Scanln(&t)
//	fmt.Printf("number of Edge : ")
//	fmt.Scanln(&t2)
//	G:= Graph{NumberOfVertex: t,NumberOfEdge: t2}
//	G.EdgeMap=make(map[int]Matrix,G.NumberOfVertex)
//	for i :=0 ;i<t2;i++ {
//		var li []int
//		fmt.Printf("give - Vertex1 , Verex2 , Weight : ")
//		b,_,_:= bufio.NewReader(os.Stdin).ReadLine()
//		sli:= strings.Split(string(b)," ")
//		for j:=0 ;j<len(sli);j++{
//			temp,_:=strconv.Atoi(sli[j])
//			li = append(li, temp)
//
//		}
//
//
//		G.AddEdge(li)
//
//
//	}
//	fmt.Println(G.EdgeList)
//	fmt.Println(G.EdgeMap)
//	fmt.Println(G.EdgeMap)
//	lix:= G.PrimsMST()
//	fmt.Println(G.EdgeList)
//	fmt.Println(lix)
//
//
//
//}
//
//type Graph struct {
//	NumberOfVertex int
//	NumberOfEdge int
//	EdgeList Matrix
//	EdgeMap map[int]Matrix
//
//}
//
//func (g *Graph)AddEdge(li []int ){
//	g.EdgeList=append(g.EdgeList,li)
//
//	g.EdgeMap[li[0]]=append(g.EdgeMap[li[0]],[]int{li[1],li[2]})
//	g.EdgeMap[li[1]]=append(g.EdgeMap[li[1]],[]int{li[0],li[2]})
//
//
//}
//
//func (g *Graph)PrimsMST()(Matrix){
//
//
//	//x:=func (s []int, index int) []int {
//	//	fmt.Println("removing : ",s,index)
//	//	return append(s[:index], s[index+1:]...)
//	//}
//
//
//	sort.Sort(g.EdgeList)
//	//added :=make([]bool,g.NumberOfEdge)
//	//cost:=0
//	fmt.Println("sorted : ",g.EdgeList)
//	ansList:=make(Matrix,0)
//	//PriorityMap := make(map[int][]int,10)
//	//PriorityMap[g.EdgeList[0][0]]=[]int{g.EdgeList[0][1],g.EdgeList[0][2]}
//	li:= make(Matrix,0)
//
//	li=append(li,g.EdgeList[0])
//	//sort.Sort(li)
//	ansList=append(ansList,li[0])
//	fmt.Println("g.EdgeMap : ",g.EdgeMap)
//	fmt.Println("li : ",li)
//	fmt.Println("g.EdgeMap[li[0][0]] : ",g.EdgeMap[li[0][0]])
//	//fmt.Println("g.EdgeMap[li[0][0]] : ",g.EdgeMap[li[0][0]])
//	fmt.Println("g.EdgeMap[li[0][1]] : ",g.EdgeMap[li[0][1]])
//	fmt.Println("")
//	for i,v := range g.EdgeMap[li[0][0]] {
//		fmt.Println("for loop 1: ",i,v)
//		li = append(li, []int{li[0][0],v[0],v[1]})
//	}
//	for i,v := range g.EdgeMap[li[0][1]] {
//		fmt.Println("for loop 2: ",i,v)
//		li = append(li, []int{li[0][1],v[0],v[1]})
//	}
//	fmt.Println("li : ",li)
//	li=append(li[:0],li[1:]...)
//	fmt.Println("li : ",li)
//
//	Isthere := func(a []int) bool {
//		for _,x := range ansList {
//			if (a[0]==x[0] && a[1]==x[1]){
//				return false
//			}
//		}
//		return true
//	}
//	Isthere2 := func(a []int) bool {
//		for _,x := range li {
//			if (a[0]==x[0] && a[1]==x[1])||(a[1]==x[0] && a[0]==x[1]){
//				return false
//			}
//		}
//
//		for _,x := range ansList {
//			if a[0]==x[0] || a[1]==x[1]||a[1]==x[0] || a[0]==x[1]{
//				return false
//			}
//		}
//		return true
//	}
//
//	for i:=0;i<g.NumberOfVertex;i++ {
//		if len(li)==0{
//			fmt.Println("break li")
//			break
//		}
//		fmt.Println("for loop 3: ",li)
//		sort.Sort(li)
//		if Isthere([]int{li[0][0],li[0][0]}) {
//			ansList=append(ansList,li[0])
//			fmt.Println("append ans : ",li[0])
//		}
//		for i,v := range g.EdgeMap[li[0][0]] {
//			fmt.Println("for loop 10: ",i,v)
//			if Isthere2([]int{li[0][0],v[0]}) {
//
//				li = append(li, []int{li[0][0],v[0],v[1]})
//			fmt.Println("append li : ",li[0][0],v[0],v[1])
//			}
//		}
//		li=append(li[:0],li[1:]...)
//
//
//	}
//
//	return ansList
//
//
//
//}
//
//// func (g *Graph)FindParent(i int)bool{ }
//type Matrix [][]int
//func (m Matrix)Len()int{return len(m)}
//func (m Matrix)Less(i,j int)bool{
//	return m[i][2]<m[j][2]
//}
//func (m Matrix)Swap(i,j int){
//	m[i],m[j]=m[j],m[i]
//}
//

//////////////////////////////////////////////////////////////////////////////////

package main

import (
"bufio"
"fmt"
"os"
	"reflect"
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
type Dictionary map[int]int
func (m Dictionary)Len()int{return len(m)}
func (m Dictionary)Less(i,j int)bool{

	return m[i][1]<m[j][1]
}
func (m Dictionary)Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}

type Node struct{
	Id int
}

type Graph struct {
	source int
	FriendList map[int][]int
}

func (g *Graph)PrimsMST()int {
	PriorityQueue := make(Dictionary,0)
	PriorityQueue[Node{g.source}]=0
	var added []bool
	for i:= 0 ;i<len(g.FriendList) ;i++{
		added=append(added,false)
	}
	min_span_tree_cost :=0
	for len(PriorityQueue)>0 {
		sort.Sort(PriorityQueue)
		node := Node{}
		for k := range PriorityQueue {
			node= k
			break
		}
		cost := PriorityQueue[node]
		delete(PriorityQueue,node)
		if added[node.Id] ==false {
			min_span_tree_cost+=cost
			added[node.Id]=true
			fmt.Println("added node : ",node.Id,"cost = ",min_span_tree_cost)
			li:=g.FriendList[node.Id]
			for len(li)>0 {
				adjnode:= li[0]
				adjcost := li[1]
				if added[adjnode]==false {
					nn:=Node{adjnode}
					PriorityQueue[node]=adjcost
				}
			}
		}
		//node := PriorityQueue
		//node := reflect.ValueOf(PriorityQueue)

	}
	return min_span_tree_cost

}