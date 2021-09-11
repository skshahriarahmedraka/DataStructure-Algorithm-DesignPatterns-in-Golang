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
	if err !=nil {
		log.Fatalln(err)
	}
}

func Linear(array []int, query int) int {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}

func main(){
	fmt.Printf("Want to give array of numbers (y/n): ")
	var (
		y string
		arr []int
		in int
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
	fmt.Print("number you want to search : ")
	fmt.Scanln(&in)

	position:=Linear(arr,in)
	fmt.Println("Position of ",in," in array : ",position)

}