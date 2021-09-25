package main

import (
	"bufio"
	"reflect"
	"strings"

	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Printf("Byte to int...\n byte : ")
	b,_,_ := bufio.NewReader(os.Stdin).ReadLine()
	i,_:= strconv.Atoi(string(b))
	fmt.Println("1.int - strconv.Atoi(string(b)) :",i," ,type : ",reflect.TypeOf(i))

	fmt.Printf("Byte to Uint64...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	ui,_:=strconv.ParseUint(string(b),10,64)
	fmt.Println("2.Uint - strconv.ParseUint(string(s),10,64) : ",ui," ,type : ",reflect.TypeOf(ui))

	fmt.Printf("Byte to string...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	s:= string(b)
	fmt.Println("3.string - string() : ",s," ,type : ",reflect.TypeOf(s))

	fmt.Printf("Byte to float64...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	f,_:=strconv.ParseFloat(string(b),64)
	fmt.Println("4.float - strconv.ParseFloat(string(b)) : ",f," ,type : ",reflect.TypeOf(f))

	fmt.Printf("Byte to Boolean...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	bb,_:=strconv.ParseBool(string(b))
	fmt.Println("5.bool - strconv.ParseBool(string(b)) : ",bb," ,type : ",reflect.TypeOf(bb))

	fmt.Printf("Byte to complex128...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	cx,_:=strconv.ParseComplex(string(b),128)
	fmt.Println("6.complex - strconv.ParseComplex(string(b),128) : ",cx," ,type : ",reflect.TypeOf(cx))

	fmt.Printf("Byte to slice...\n byte : ")
	b,_,_ = bufio.NewReader(os.Stdin).ReadLine()
	sli:= strings.Split(string(b)," ")
	fmt.Println("7.slice - strings.Split(string(b),\" \") : ",sli," ,type : ",reflect.TypeOf(sli))
}
