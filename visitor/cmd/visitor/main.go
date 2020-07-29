package main

import (
	"fmt"
	"github.com/wb_go/visitor/pkg/customintslice"
	"github.com/wb_go/visitor/pkg/visitor"
)

//TODO

func main() {
	a := []int{1, 2, 7, 3, 22}
	//b := []float64{1.1, 2.2, 3.4343, 5.5}
	//c := []string{"one", "two", "three"}

	//fl := customfloatslice.NewCustomFloatSlice(b)
	//str := customstrslice.NewCustomStrSlice(c)
	in := customintslice.NewCustomIntSlice(a)

	arrFl := visitor.NewArr(5)
	//
	//fmt.Println(fl.Accept(arrFl))
	//fmt.Printf("%v", b)
	//
	//
	//str.Accept(arrFl)
	//fmt.Printf("%v", c)

	in.Accept(arrFl)
	fmt.Printf("%v", in)

}
