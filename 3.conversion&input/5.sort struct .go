


package main

import (
	"fmt"
	"sort"
)

type people struct {
	Name string
	Roll int
	Money float64
}

type StructSort []people

func (a StructSort)Len() int {return len(a)}
func (a StructSort)Less(i,j  int) bool {return a[i].Name <a[j].Name}
func (a StructSort) Swap(i,j int) {a[i],a[j]=a[j],a[i]}




func main(){
	li := []people{
		{"raka",180107,123},
		{"raka",180109,1223},
		{"raka",180109,1223},
		{"raka",180109,1},
		{"ahmed",101,1222},
		{"ad",103,12},
	}
	sort.Sort(StructSort(li))
	fmt.Println(li)
}