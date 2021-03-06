package main

import (
	"fmt"

	"github.com/wb_go/facade/pkg/changer"
	"github.com/wb_go/facade/pkg/customfile"
	"github.com/wb_go/facade/pkg/facade"
)

const (
	errorReceiver = "Error occurred while Receive method called"
)

var (
	list          = []string{"test1.old", "test2.old", "test3.old"}
	refreshedList string
)

func main() {
	for i := range list {
		change := changer.NewChanger(list[i])
		cfile := customfile.NewCustomFile(refreshedList)
		rec := facade.NewConverter(cfile, change)
		res, err := rec.Receive(list[i], ".lal")
		if err != nil {
			fmt.Println(errorReceiver, err)
			return
		}
		fmt.Println(res)
	}
}
