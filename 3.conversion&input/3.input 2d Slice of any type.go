package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main (){
	fmt.Println("### 2d slice of any type ####")
	fmt.Println("n x m  slice ...")
	fmt.Printf("value n : ")
	var n,m int
	fmt.Scanln(&n)
	fmt.Printf("value m : ")
	fmt.Scanln(&m)
	fmt.Println("give the values according to row : ")
	sli := make([][]interface{},0)
	for i:=0;i<n;i++{
		li:=make([]interface{},0)
		for j:=0;j<m;j++{
			x,_,_ :=bufio.NewReader(os.Stdin).ReadLine()
			ans:=Anytype(string(x))
			li=append(li,ans)
		}
		sli=append(sli,li)

	}
	fmt.Println(sli)
	fmt.Printf("each value type : %T\n",sli[0][0])

}


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