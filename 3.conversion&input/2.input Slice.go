package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var s string


	fmt.Println("### bufio.NewScanner(os.Stdin) .Scan()  .Text()###########")
	x := bufio.NewScanner(os.Stdin)
	x.Scan()
	s=x.Text()
	//fmt.Println("input string : ",s)
	//fmt.Printf("type : %T ,value : %v",s,s)
	li:= strings.Split(string(s)," ")
	var  li2 []int
	for i:=0;i<len(li) ;i++{
		xx,_:=strconv.Atoi(li[i])
		li2=append(li2,xx)
	}
	fmt.Printf("type : %T ,value : %v\n",li2,li2)



	fmt.Println("### bufio.NewReader(os.Stdin).ReadString('\\n') strings.TrimSuffix(Name,\"\\n\")###### ")
	Name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s=strings.TrimSuffix(Name,"\n")
	//fmt.Println(Name)
	//fmt.Printf("type : %T ,value : %v",s,s)
	li= strings.Split(s," ")
	var  li3 []int
	for i:=0;i<len(li) ;i++{
		xx,_:=strconv.Atoi(li[i])
		li3=append(li3,xx)
	}
	fmt.Printf("type : %T ,value : %v\n",li3,li3)




	fmt.Println("### bufio.NewReader(os.Stdin).ReadLine() string(N) ######")
	N, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	s=string(N)
	li= strings.Split(s," ")
	var  li4 []int
	for i:=0;i<len(li) ;i++{
		xx,_:=strconv.Atoi(li[i])
		li4=append(li4,xx)
	}
	fmt.Printf("type : %T ,value : %v\n",li4,li4)


}