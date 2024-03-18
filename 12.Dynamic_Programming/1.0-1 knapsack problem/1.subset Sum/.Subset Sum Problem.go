package main

import (
	// "bufio"
	// "fmt"
	// "os"
	// "strconv"
	// "strings"
)

// func main (){
// 	fmt.Printf("give a List of int : ")
// 	N, _, _ := bufio.NewReader(os.Stdin).ReadLine()
// 	s:=string(N)
// 	li:= strings.Split(s," ")
// 	var  li4 []int
// 	var sum int
// 	for i:=0;i<len(li) ;i++{
// 		xx,_:=strconv.Atoi(li[i])
// 		li4=append(li4,xx)
// 	}

// 	fmt.Printf("type : %T ,value : %v\n",li4,li4)
// 	fmt.Printf("sub set Sum : ")
// 	fmt.Scanln(&sum)
// 	 if IsSubSetSum(li4,len(li),sum){
// 		 fmt.Println(sum,"sum available")
// 	 }else{
// 		 fmt.Println(sum, "sum not available")
// 	 }
// }


func IsSubSetSum(li []int,n int,sum int) bool{
	if sum ==0 {
		return true
	}
	if n==0 {
		return false
	}
	if li[n-1]>sum {
		return IsSubSetSum(li,n-1,sum)
	}
	return IsSubSetSum(li,n-1,sum) || IsSubSetSum(li,n-1,sum-li[n-1])
}