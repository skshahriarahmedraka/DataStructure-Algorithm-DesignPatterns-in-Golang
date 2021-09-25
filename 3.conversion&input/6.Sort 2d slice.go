package main

import (
	"bufio"
	//"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Matrix [][]int

func (m Matrix)Len()int {return  len(m)}
func (m Matrix)Less(i,j int) bool{
	//for x:= range m[i] {
	//	if m[i][x] == m[j][x] {
	//		continue
	//	}
	//	return m[i][x]<m[j][x]
	//}
	//return false
	return m[i][1]<m[j][1]
}
func (m Matrix)Swap(i,j int ){
	m[i],m[j]=m[j],m[i]
}


func main (){
	fmt.Println("### 2d slice of any type ####")
	fmt.Println("n x m  slice ...")
	fmt.Printf("value n : ")
	var n,m int
	fmt.Scanln(&n)
	fmt.Printf("value m : ")
	fmt.Scanln(&m)
	fmt.Println("give the values according to row : ")
	sli := make(Matrix,0)
	for i:=0;i<n;i++{
		li:=make([]int,0)
		for j:=0;j<m;j++{
			x,_,_ :=bufio.NewReader(os.Stdin).ReadLine()
			ans,_:=strconv.Atoi(string(x))
			li=append(li,ans)
		}
		sli=append(sli,li)
	}
	fmt.Println(sli)
	sort.Sort(sli)
	fmt.Println(sli)

}
