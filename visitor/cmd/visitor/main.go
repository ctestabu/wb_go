package main

import (
	"fmt"

	"github.com/wb_go/visitor/pkg/customfloatslice"
	"github.com/wb_go/visitor/pkg/customintslice"
	"github.com/wb_go/visitor/pkg/customstrslice"
	"github.com/wb_go/visitor/pkg/visitor"
)

const (
	errorCustomSlice = "Error with custom slice"
)

func main() {
	a := []int{1, 2, 7, 3, 22}
	b := []float64{1.1, 2.2, 3.4343, 5.5}
	c := []string{"one", "two", "three"}
	test := 4

	in1 := customintslice.NewCustomIntSlice(a)
	in2 := customfloatslice.NewCustomFloatSlice(b)
	in3 := customstrslice.NewCustomStrSlice(c)

	val1 := visitor.NewArr(test)
	err := in1.Accept(val1)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in1)

	val2 := visitor.NewArr(test)
	err = in2.Accept(val2)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in2)

	val3 := visitor.NewArr(test)
	err = in3.Accept(val3)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in3)

}
