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
func MargeSort(arr []int,n int) {
	if n>1{
		mid := int(n/2)
		L:= arr[:mid]
		R:= arr[mid:]

		MargeSort(L,len(L))
		MargeSort(R,len(R))

		i:=0;j:=0;k:=0
		for i<len(L) && j<len(R){
			if L[i] <R[j]{
				arr[k] = L[i]
				i++
			}else {
				arr[k] = R[j]
				j++
			}
			k++
		}
		for i<len(L){
			arr[k]=L[i]
			i++
			k++
		}
		for j <len(R){
			arr[k] = R[j]
			j++
			k++
		}
	}
}
func main (){
	fmt.Println("*******Marge Sort******")
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

	MargeSort(arr,len(arr))
	fmt.Println("Marge Sorted array : ",arr)
}
