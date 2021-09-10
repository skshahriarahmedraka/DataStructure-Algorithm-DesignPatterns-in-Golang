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
func InsertionSort(arr *[]int){
	var (
		i,j,k int
	)
	for i=1 ;i <len(*arr);i++{
		k=(*arr)[i]

		j=i-1
		for j>=0 && k< (*arr)[j]{
			(*arr)[j+1]=(*arr)[j]
			j--
		}
		(*arr)[j+1]=k

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

	InsertionSort(&arr)
	fmt.Println("insertion  Sorted array : ",arr)
}
