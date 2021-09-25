package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func Anytype(s string )interface{}{

	i,err :=strconv.Atoi(s)
	if err==nil{
		return i
	}
	//i32,err := strconv.ParseInt(s,10,32)
	//fmt.Println(i32," ",err)
	//if err==nil{
	//	return i32
	//}
	i64,err := strconv.ParseInt(s,10,64)
	if err==nil{
		return i64
	}
	ui64,err := strconv.ParseUint(s,10,64)
	if err==nil{
		return ui64
	}
	f32,err := strconv.ParseFloat(s,32)
	// fmt.Println(f32," ",err)
	if err==nil{
		return f32
	}
	f64,err := strconv.ParseFloat(s,64)
	if err==nil{
		return f64
	}

	com64,err := strconv.ParseComplex(s,64)
	if err==nil{
		return com64
	}
	com128,err := strconv.ParseComplex(s,128)
	if err==nil{
		return com128
	}
	b,err := strconv.ParseBool(s)// "1", "t", "T", "true", "TRUE", "True"//"0", "f", "F", "false", "FALSE", "False"
	if err==nil{
		return b
	}
	return s
}

func main(){
	fmt.Printf("give any type of input (1 input): ")
	x,_,_ :=bufio.NewReader(os.Stdin).ReadLine()

	ans:=Anytype(string(x))

	fmt.Printf("type : %T ,value : %v\n",ans,ans)

}