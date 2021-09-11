package main

import "fmt"
func RecursiveGCD(a,b uint64)uint64{
	if b==0 {
		return a
	}
	return RecursiveGCD(b,a%b)
}
func main(){
	fmt.Println("********GCD**********")
	var a,b,ans uint64
	fmt.Printf("Give First number : ")
	fmt.Scanln(&a)
	fmt.Printf("Give 2nd number : ")
	fmt.Scanln(&b)
	if a>=b{
		ans=RecursiveGCD(a,b)
	}else{
		ans=RecursiveGCD(b,a)
	}
	LCM := (a*b)/ans
	fmt.Println("LCM : ",LCM)
}
