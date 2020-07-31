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

	visit := visitor.NewArr(test)
	in1 := customintslice.NewCustomIntSlice(a)
	in2 := customfloatslice.NewCustomFloatSlice(b)
	in3 := customstrslice.NewCustomStrSlice(c)

	err := in1.Accept(visit)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in1)

	err = in2.Accept(visit)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in2)

	err = in3.Accept(visit)
	if err != nil {
		fmt.Println(errorCustomSlice, err)
	}
	fmt.Println(in3)
}
