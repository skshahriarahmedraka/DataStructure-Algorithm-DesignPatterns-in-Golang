package main

import "fmt"

func IterativePower(n uint64, power uint64) uint64 {
	var res uint64 = 1
	for power > 0 {
		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1
		n *= n
	}
	return res
}

func main()  {
	fmt.Println("****************power***************")
	var a,x uint64
	fmt.Println("format : (a^x) ")
	fmt.Printf("a (base) : ")
	fmt.Scanln(&a)
	fmt.Printf("x (power): ")
	fmt.Scanln(&x)
	n:=IterativePower(a,x)
	fmt.Println("answer : ",n)
}
