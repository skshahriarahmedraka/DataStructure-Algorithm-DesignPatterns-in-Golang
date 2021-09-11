/**
 * @Author: sk shahriar ahmed raka <ssar>
 * @Date:   01-Jan-1970
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: r2.go
 * @Last modified by:   ssar
 * @Last modified time: 27-Mar-2021
 */



package main

import (

  "fmt"
)
type geometry interface {
  area() float64
  area2() float64
}

type squire struct {
  height float64
  width float64
}
func (r squire) area() float64 {
  return r.height * r.width
}
func (r squire) area2() float64 {
  return r.height * r.width *2
}

func f1 (g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
}

func main(){
  r1 := squire{height : 3,width :4}
  // fmt.Println(type(r1))
  f1(r1)


}
