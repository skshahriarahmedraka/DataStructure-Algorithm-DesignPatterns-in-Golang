/**
 * @Author: sk shahriar ahmed raka <ssar>
 * @Date:   01-Jan-1970
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: r1.go
 * @Last modified by:   ssar
 * @Last modified time: 27-Mar-2021
 */



package main

import (
    "fmt"
    "math"
)

// interface for shapes:
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

// implement all methods for interface:
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main(){

    r := rect{width: 5, height: 10}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
