/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */



package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



func CheckError(err error){
	if err != nil{
		log.Fatalln(err)
	}
}
func heapify(arr *[]int,n,i int){
	largest :=i
	l:= 2*i+1
	r:=2*i+2

	if l<n && (*arr) {

	}
}
func Sort(arr *[]int ) {

	n := len(*arr)
	for i:=int(n/2)-1 ;i>=0;i--{
		heapify(arr,n,i)
	}
	for i:=n-1 ;i>=0 ;i++{
		heapify(arr,i,0)
	}
}
func main (){
	fmt.Printf("Want to give array of numbers (y/n): ")
	var (
		y string
		arr []int
	)
	fmt.Scanln(&y)
	if y=="y"{
		fmt.Printf("array of numbers : ")
		b,_,err:=bufio.NewReader(os.Stdin).ReadLine()
		CheckError(err)
		s:= string(b)
		if len(s)!=0 {
			s2:= strings.Split(s," ")
			for i:=0; i<len(s2);i++{
				a,err:=strconv.Atoi(s2[i])
				CheckError(err)
				arr=append(arr,a)
			}
		}
	}else {
		arr =append(arr, 2,9,3,6,2,4,4,0)
		fmt.Println("Sample array : ",arr)
	}

	Sort(&arr,len(arr))
	fmt.Println("RecursiveInsertionSort  Sorted array : ",arr)
}
