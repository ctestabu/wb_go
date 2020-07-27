//TODO: Порядок работы программы - Приходит строка с названием файла --- изменяется на нужный формат
//TODO --- заполняется список измененных файлов --- и создается файл с нужным название. Не забыть логи удалить потом файлы
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
		change := changer.New(list[i])
		cfile := customfile.New(refreshedList)
		rec := facade.NewConverter(cfile, change)
		res, err := rec.Receive(list[i], ".lal")
		if err != nil {
			fmt.Println(errorReceiver)
			return
		}
		fmt.Println(res)
	}

}
