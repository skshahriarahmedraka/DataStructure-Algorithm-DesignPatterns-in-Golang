/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	fmt.Println("########### 0-1 knapsack ###########")
	fmt.Printf("give the busket capacity value : ")
	var MaxCapacity int
	fmt.Scanln(&MaxCapacity)
	fmt.Printf("Give the product Value List : ")
	b,_,_:=bufio.NewReader(os.Stdin).ReadLine()
	ValueList:= IntList(strings.Split(string(b)," "))
	NumberOfValue :=len(ValueList)

	fmt.Printf("Give the product Weight List : ")
	b,_,_=bufio.NewReader(os.Stdin).ReadLine()
	WeightList:= IntList(strings.Split(string(b)," "))



	Knap:= KnapSack(MaxCapacity,WeightList,ValueList,NumberOfValue)
	fmt.Println(Knap)
}
func IntList(li []string)([]int){
	var li2 []int
	for i:=0;i<len(li);i++ {
		s,_:=strconv.Atoi(li[i])
		li2=append(li2,s)

	}
	return li2
}


func KnapSack(MaxCapacity int,WeightList []int,ValueList []int,NumberOfValue int)int{
	if MaxCapacity==0 || NumberOfValue==0 {
		return 0
	}
	if WeightList[NumberOfValue-1]>MaxCapacity {
		return KnapSack(MaxCapacity,WeightList,ValueList,NumberOfValue-1)
	}else{
		return Max(ValueList[NumberOfValue-1]+KnapSack(MaxCapacity-WeightList[NumberOfValue-1],WeightList,ValueList,NumberOfValue-1),
			KnapSack(MaxCapacity,WeightList,ValueList,NumberOfValue-1))
	}


}

func Max(a,b int ) int {
	if a>=b {
		return a
	}
	return b

}
