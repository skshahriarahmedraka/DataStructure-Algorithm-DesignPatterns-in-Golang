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
func Partition(start ,end int , arr *[]int)int{
	pivotIndex:= start
	pivot :=(*arr)[pivotIndex]

	for start<end{
		for start<len(*arr) && (*arr)[start] <= pivot {
			start+=1
		}
		for (*arr)[end] > pivot{
			end-=1
		}
		if start<end {
			(*arr)[start],(*arr)[end]=(*arr)[end],(*arr)[start]
		}
	}
	(*arr)[end],(*arr)[pivotIndex]=(*arr)[pivotIndex],(*arr)[end]
	return end
}
func QuickSort(start,end int ,arr *[]int ) {
	if (start <end){
		p:= Partition(start,end,arr)
		QuickSort(start,p-1,arr)
		QuickSort(p+1,end,arr)
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

	QuickSort(0,len(arr)-1,&arr)
	fmt.Println("RecursiveInsertionSort  Sorted array : ",arr)
}
