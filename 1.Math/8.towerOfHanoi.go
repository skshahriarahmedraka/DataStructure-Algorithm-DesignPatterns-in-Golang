package main


import (
	"fmt"
)

func towerOfHanoi(n int64,from,to,aux string ){
	if n==1 {
		fmt.Println("move disk 1 from ",from," to ",to)
		return
	}
	towerOfHanoi(n-1 ,from,aux,to)
	fmt.Println("move disk ",1," from ",from," to ",to)
	towerOfHanoi(n-1,aux,to,from)

}


func main(){
	fmt.Println("Tower of Hanoi ")
	var (
		A string ="A"
		B string ="B"
		C string ="C"
		N int64 =3
	)
	towerOfHanoi(N,A,B,C)
}