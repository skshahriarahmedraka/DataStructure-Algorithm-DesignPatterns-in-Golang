package main

import (
	"errors"
	"fmt"
)


func Exponentiation(base ,exponent,mod uint64)(uint64,error){
	if mod ==1 || base%mod ==0 {
		return 0,nil
	}
	if exponent < 0 {
		return 0 , errors.New("negative exponent ")
	}
	if base==0 || mod==0 || exponent==0 {
		return 0 ,errors.New("divide 0 error")
	}
	var (
		e,i,x,r,ans uint64
		li []uint64
	)
	for e=1 ;e<=exponent;e=e*2{
		li=append(li, e)
	}
	e=exponent
	r=base%mod
	fmt.Println(li,len(li))
	for i=uint64(len(li)-1) ;i>0 ;i--{
		if e-li[i]>=0{
			fmt.Println(li,len(li))

			x=x+(r*li[i])%mod
		}

	}
	ans =x%mod
	return ans,nil


}

func main (){
	fmt.Println("****************Modular Exponentiation***************")
	var a,x,b,ans uint64
	fmt.Println("format : (a^x) % b ")
	fmt.Printf("a : ")
	fmt.Scanln(&a)
	fmt.Printf("x : ")
	fmt.Scanln(&x)
	fmt.Printf("b : ")
	fmt.Scanln(&b)
	ans,_=Exponentiation(a,x,b)

	fmt.Println("answer : ",ans)
}
