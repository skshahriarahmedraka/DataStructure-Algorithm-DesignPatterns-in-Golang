package main

import "fmt"


func RecursiveGCD(a,b uint64)uint64{
	if b==0 {
		return a
	}
		return RecursiveGCD(b,a%b)
}
func IterativeGCD(a,b uint64)uint64{
	if a>=b {a,b=a,b}else{a,b=b,a}
	for b>0{
		a,b=b,a%b
	}
	return a
}
func main (){
	fmt.Println("********GCD**********")
	var a,b,ans uint64
	fmt.Printf("Give First number : ")
	fmt.Scanln(&a)
	fmt.Printf("Give 2nd number : ")
	fmt.Scanln(&b)
	if a>=b {a,b=a,b}else{a,b=b,a}
	//ans2:=RecursiveGCD(a,b)


	ans=IterativeGCD(a,b)
	fmt.Println("GCD : ",ans)
}
