package main

import "sort"



func main(){
	G := Graph{}
	G.V[0]=[]int{0,1,10}
	G.V[1]=[]int{0,2,6}
	G.V[2]=[]int{0,3,5}
	G.V[3]=[]int{1,3,15}
	G.V[4]=[]int{2,3,4}

	KruskalMST(G)

}

type Graph struct {
	V [][]int
}




func KruskalMST(G Graph){

}
