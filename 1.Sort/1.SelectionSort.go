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
func SelectionSort(arr []int)([]int){
	n:= len(arr)
	minIndex:=0
	for i:=0 ; i<n ;i++{
		minIndex=i
		for j:=i+1 ;j<n;j++{
			if arr[j]<arr[minIndex] {
				minIndex=j
			}
		}
		arr[i],arr[minIndex]=arr[minIndex],arr[i]
	}
	return arr
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

	sortedArray:=SelectionSort(arr)
	fmt.Println("Selection Sorted array : ",sortedArray)
}
