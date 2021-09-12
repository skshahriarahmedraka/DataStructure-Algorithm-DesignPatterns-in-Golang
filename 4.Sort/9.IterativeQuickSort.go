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
func IterativePartition(start ,end int , arr *[]int)int{
	i:=(start -1)
	pivot := (*arr)[end]

	for j:= start ; j<end ;j++{
		if (*arr)[j] <= pivot {
			i+=1
			(*arr)[i],(*arr)[j]=(*arr)[j],(*arr)[i]
		}
	}
	(*arr)[i+1],(*arr)[end]=(*arr)[end],(*arr)[i+1]

	return (i+1)

}
func IterativeQuickSort(start,end int ,arr *[]int ) {
	if (start <end){
		p:= IterativePartition(start,end,arr)
		IterativeQuickSort(start,p-1,arr)
		IterativeQuickSort(p+1,end,arr)
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

	IterativeQuickSort(0,len(arr)-1,&arr)
	fmt.Println("RecursiveInsertionSort  Sorted array : ",arr)
}
