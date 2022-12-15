package  main

import (
	"fmt"
	"math/rand"
	"time"
	
)

func main(){
	RandomIntWithOutSeed()
	RandomInt()
	RandomIntn()

	RandomFloatWithOutSeed()
	RandomFloar64()

	RandomPerm()// generate a list of random int number
}
func RandomIntWithOutSeed(){
	s:= rand.Int()
	fmt.Printf("##  Random Integer without Seed \n{rand.Int()} : %d \n\n",s)
}
func RandomInt(){
	rand.Seed(time.Now().UnixNano())
	s:= rand.Int()
	fmt.Printf("##  Random Integer with Seed \n{rand.Int()} : %d \n\n",s)
}
func RandomIntn(){
	rand.Seed(time.Now().UnixNano())
	min:=10
	max:=20
	//s:= rand.Intn(100)
	s:= rand.Intn(max-min)+min
	fmt.Printf("##  Random Integer with Seed with Number Range\n{rand.Intn(max-min)+min} : %d\n\n",s)
}


func RandomFloatWithOutSeed(){
	s:=rand.Float64()
	fmt.Printf("##  Random Float without Seed \n{rand.Float64()} : %f\n\n",s)
}
func RandomFloar64(){
	rand.Seed(time.Now().UnixNano())
	s:=rand.Float64()
	fmt.Printf("##  Random Float with Seed \n{rand.Float64()} : %f\n\n",s)
}


func RandomPerm(){
	rand.Seed(time.Now().UnixNano())
	s:=rand.Perm(10)
	fmt.Printf("## Random List of Integer Number \n{rand.Perm(n) : %v\n\n",s)
}