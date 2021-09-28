package  main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("give number an array of int : ")

	//fmt.Println("### bufio.NewReader(os.Stdin).ReadLine() string(N) ######")
	N, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	s:=string(N)
	li:= strings.Split(s," ")
	var  li4 []int
	for i:=0;i<len(li) ;i++{
		xx,_:=strconv.Atoi(li[i])
		li4=append(li4,xx)
	}
	var num int
	fmt.Printf("number you want to search : ")
	fmt.Scanln(&num)
	ans:=BinarySearch(li4,0,len(li4)-1,num)
	if ans!= -1 {
		fmt.Println("number is found index : ",ans)
	}else {
		fmt.Println("sorry not found")
	}

}

func BinarySearch(arr []int,l,r,n int)int {
	for l<=r {
		mid:= l+(r-l)/2
		if arr[mid]==n {
			return mid
		}else if arr[mid]<n {
			l=mid +1
		}else {
			r=mid-1
		}
	}
	return -1
}