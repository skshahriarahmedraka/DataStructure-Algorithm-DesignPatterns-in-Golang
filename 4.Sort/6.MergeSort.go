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


func marge(a []int,b []int) []int {
	r:= make ([]int, len(a)+len(b))
	i,j:=0,0
	for ; i<len(a) && j<len(b);{
		if a[i] <= b[j]{
			r[i+j]=a[i]
			i++
		}else{
			r[i+j]=b[j]
			j++
		}
	}
	for i<len(a){
		r[i+j]=a[i]
		i++
	}
	for j<len(b) {
		r[i+j]=b[j]
		j++
	}
	return r
}

func MargeSort(arr []int) []int {
	if len(arr)<2 {
		return arr
	}

	m := int(len(arr)/2)
	a := MargeSort(arr[:m])
	b:= MargeSort(arr[m:])
	return marge(a,b)
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

	arr2:=MargeSort(arr)
	fmt.Println("Marge Sorted array : ",arr2)
}
