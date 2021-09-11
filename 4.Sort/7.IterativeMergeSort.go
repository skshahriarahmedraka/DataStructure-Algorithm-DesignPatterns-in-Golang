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
func merge2(arr []int ,L1 int,R1 int,L2 int ,R2 int )[]int{
	var temp []int
	for ii:=0 ;ii<len(arr);ii++{
		temp=append(temp,0)
	}
	index :=0
	for L1<=R1 && L2 <=R2 {
		if arr[L1]<=arr[L2]{
			temp[index]=arr[L1]
			index++
			L1=L1+1
		}else {
			temp[index]=arr[L2]
			index++
			L2=L2+1
		}
	}
	for L1<=R1 {
		temp[index]=arr[L1]
		index++
		L1++
	}
	for L2<=R2 {
		temp[index]=arr[L2]
		index++
		L2++
	}
	return temp
}

func IterativeMargeSort(arr []int,n int) []int {
	//if len(arr)<2 {
	//	return arr
	//}
	//
	//m := int(len(arr)/2)
	//a := MargeSort(arr[:m])
	//b:= MargeSort(arr[m:])
	//return marge(a,b)


	l:=1
	for l <n {
		for i:=0;i<n; {
			L1 :=i
			R1 := i+l -1
			L2 := i+l
			R2 := i+2*l-1
			if L2 >= n {
				break
			}
			if R2>=n {
				R2 = n-1
			}
			temp := merge2(arr,L1,R1,L2,R2)
			for j:=0 ;j <= R2-L1+1 ;j++ {
				arr[i+j]=temp[j]
			}
			i=i+2*l
		}
		l=2*l
	}
	return arr

}

func main (){
	fmt.Println("******* Iterative Marge Sort******")
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

	arr2:=IterativeMargeSort(arr,len(arr))
	fmt.Println("Marge Sorted array : ",arr2)
}
