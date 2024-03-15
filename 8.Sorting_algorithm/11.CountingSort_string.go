package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
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
	fmt.Printf("Want to give string (y/n): ")
	var (
		y string
		//arr []string
		s2 string
	)
	fmt.Scanln(&y)
	if y=="y"{
		fmt.Printf("string : ")
		b,_,err:=bufio.NewReader(os.Stdin).ReadLine()
		CheckError(err)
		s:= string(b)
		if len(s)!=0 {
			//s2:= strings.Split(s," ")
			s2= strings.Replace(s, " ", "", -1)

		}
	}else {
		s2="sk shahriar ahmed raka"
		fmt.Println("Sample array : ",s2)
	}

	ss:=CountingSort(&s2,len(s2))
	fmt.Println("CountingSort  Sorted array : ",ss)
}
