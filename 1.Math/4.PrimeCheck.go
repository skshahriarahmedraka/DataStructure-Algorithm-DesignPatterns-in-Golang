package main

import "fmt"

func NaiveApproach(n uint64) bool {
	if n < 2 {return false}
	var i uint64= 2
	for ; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func PairApproach(n uint64) bool {
	if n < 2 {return false}
	var i uint64= 2
	for ; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println("****************Prime check***************")
	var a uint64
	fmt.Printf("give Number : ")
	fmt.Scanln(&a)
	b1:=NaiveApproach(a)
	//b2:=PairApproach(a)
	if b1{
		fmt.Println(" Prime Number ")
	}else{
		fmt.Println("Not Prime Number ")

	}
}