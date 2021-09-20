package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	//"strconv"
)

func main(){
	var s,p string
	fmt.Println("********** Rabin-Karp Algorithm *************")
	fmt.Printf("\ngive the text : ")
	a,_,_:= bufio.NewReader(os.Stdin).ReadLine()
	s=string(a)
	fmt.Printf("\nGive pattern : ")
	a,_,_= bufio.NewReader(os.Stdin).ReadLine()
	p=string(a)

	li:=RabinKarp(s,p)
	fmt.Println(li)
}



func RabinKarp(s string,p string)[]uint64{

	slen:= uint64(len(s))
	plen := uint64(len(p))
	var num,num2 uint64
	num=0
	num2=0
	for i:= uint64(0);i<plen ;i++{
		s:=uint64(s[i])
		num+= s*uint64((math.Pow(100,float64(plen-1-i))))
	fmt.Println(num)
	}

	for i:= uint64(0);i<plen ;i++{
		s:=uint64(p[i])
		num2+= s*uint64(math.Pow(100,float64(plen-1-i)) )
	fmt.Println(num2)
	}
	var li []uint64
	for i:= uint64(0);i<slen-plen;i++{
		fmt.Println("nums : ",num," ",num2)
		if num == num2 {
			li=append(li,i)
			num-= uint64(s[i])*uint64(math.Pow(100,float64(plen-1)))
			num*=100
			num+=uint64(s[i+plen])
		}else {
			num-= uint64(s[i])*uint64(math.Pow(100,float64(plen-1)))
			num*=100
			num+=uint64(s[i+plen])
		}

	}
	return li


}