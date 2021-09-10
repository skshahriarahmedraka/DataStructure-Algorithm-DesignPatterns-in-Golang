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

func BubbleSort(arr []int)([]int){
	n:= len(arr)

	for i:=0 ;i <n;i++{
		for j:=0 ;j<n-i-1;j++{
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	return arr

}

func main (){
	fmt.Printf("want to give array of numbers (y/n): ")
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

	sortedArray:=BubbleSort(arr)
	fmt.Println("Bubble Sorted array : ",sortedArray)
}
