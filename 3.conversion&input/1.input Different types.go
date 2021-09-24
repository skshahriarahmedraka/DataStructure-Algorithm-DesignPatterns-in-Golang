package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	var (
		b bool

		i int
		i8 int8
		i16 int16
		i32 int32
		i64 int64

		ui uint
		ui8 uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64

		f32 float32
		f64 float64

		c64 complex64
		c128 complex128

		s string

		p *string

		//cn chan int
		arr [5]string

	)
		m := make(map[int]string)
	fmt.Printf("boolean (true/false): ")
	fmt.Scanf("%s",&s)
	b,_=strconv.ParseBool(s)
	fmt.Printf("type : %T ,value : %v\n",b,b)

	fmt.Printf("int : ")
	fmt.Scanf("%d",&i)
	fmt.Printf("type : %T ,value : %d \n",i,i)

	fmt.Printf("int8 : ")
	fmt.Scanf("%d",&i8)
	fmt.Printf("type : %T ,value : %d \n",i8,i8)

	fmt.Printf("int16 : ")
	fmt.Scanf("%d",&i16)
	fmt.Printf("type : %T ,value : %d \n",i16,i16)

	fmt.Printf("int32 : ")
	fmt.Scanf("%d",&i32)
	fmt.Printf("type : %T ,value : %d \n",i32,i32)

	fmt.Printf("int64 : ")
	fmt.Scanf("%d",&i64)
	fmt.Printf("type : %T ,value : %d \n",i64,i64)

	fmt.Printf("uint : ")
	fmt.Scanf("%d",&ui)
	fmt.Printf("type : %T ,value : %d \n",ui,ui)

	fmt.Printf("uint8 : ")
	fmt.Scanf("%d",&ui8)
	fmt.Printf("type : %T ,value : %d \n",ui8,ui8)
	fmt.Printf("uint16 : ")
	fmt.Scanf("%d",&ui16)
	fmt.Printf("type : %T ,value : %d \n",ui16,ui16)
	fmt.Printf("uint32 : ")
	fmt.Scanf("%d",&ui32)
	fmt.Printf("type : %T ,value : %d \n",ui32,ui32)

	fmt.Printf("uint64 : ")
	fmt.Scanf("%d",&ui64)
	fmt.Printf("type : %T ,value : %d \n",ui64,ui64)

	//%f     default width, default precision
	//%9f    width 9, default precision
	//%.2f   default width, precision 2
	//%9.2f  width 9, precision 2
	//%9.f   width 9, precision 0

	fmt.Printf("float32 : ")
	fmt.Scanf("%f",&f32)
	fmt.Printf("type : %T ,value : %f \n",f32,f32)

	fmt.Printf("float64 : ")
	fmt.Scanf("%f",&f64)
	fmt.Printf("type : %T ,value : %f \n",f64,f64)


	fmt.Printf("complex64 (real): ")
	fmt.Scanf("%f",&f32)
	x:=float32(0)
	fmt.Printf("(imaginary): ")
	fmt.Scanf("%f",&x)
	c64=complex(f32,x)
	fmt.Printf("type : %T ,value : %f \n",c64,c64)
	fmt.Println("real Part : ",real(c64)," imaginary part : ",imag(c64)," \n ")


	fmt.Printf("complex128 (real): ")
	fmt.Scanf("%f",&f64)
	x2:=float64(0)
	fmt.Printf("(imaginary): ")
	fmt.Scanf("%f",&x)
	c128=complex(f64,x2)
	fmt.Printf("type : %T ,value : %f \n\n",c128,c128)
	fmt.Println("real Part : ",real(c128)," imaginary part : ",imag(c128)," \n ")

	p=&s
	fmt.Printf("type : %T , pointer adress : %p\n",p,p)

	fmt.Printf(" array (len 5) : ")
	s1,_,_:=bufio.NewReader(os.Stdin).ReadLine()
	s= string(s1)
	xx := strings.Split(s," ")
	copy(arr[:] ,xx[0:5])
	fmt.Printf("type : %T ",arr)
	fmt.Println("array : ",arr)

	fmt.Printf("number of value in Map : ")
	fmt.Scanf("%d",&i)
	for j,k:=0,"";i>0;i--{
		fmt.Printf("key (int): ")
		fmt.Scanf("%d",&j)
		fmt.Printf("value (string): ")
		fmt.Scanf("%s",&k)
		m[j]=k
	}
	fmt.Printf("type %T , value : %v\n",m,m)


}
