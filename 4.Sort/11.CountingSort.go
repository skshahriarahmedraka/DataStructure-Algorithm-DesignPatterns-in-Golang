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
func CountingSort(arr *string, n int ) string {
	var lenArr,count []int
	var ans []string
	for i:=0;i<n;i++{
		lenArr=append(lenArr,0)
	}
	for i:=0;i<257;i++{
		count=append(count,0)
	}

	for i:=0;i<n;i++{
		ans=append(ans,"")
	}

	for i:=0 ; i<n ;i++{
		count[int((*arr)[i])]+=1
	}
	var s []string
	for i:=1;i<257;i++{
		if count[i]>0{
			s=append(s,strings.Repeat(string(i),count[i]))
		}
	}
	ss:= strings.Join(s,"")

return ss


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
	arr2:="bcdar"
	ss:=CountingSort(&arr2,len(arr2))
	fmt.Println("RecursiveInsertionSort  Sorted array : ",ss)
}
