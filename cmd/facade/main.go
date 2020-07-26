//TODO: Порядок работы программы - Приходит строка с названием файла --- изменяется на нужный формат
//TODO --- заполняется список измененных файлов --- и создается файл с нужным название. Не забыть логи удалить потом файлы
package main

import (
	"facade/pkg/changer"
	"facade/pkg/customFile"
	"facade/pkg/facade"
	"fmt"
)

var (
	list = []string{"test1.old", "test2.old", "test3.old"}
	refreshedList = make([]string, 3)
)

func main() {
	change := changer.New(list)
	cfile := customFile.New(refreshedList)

	rec := facade.NewConverter(cfile, change)

	res, err := rec.Receive(list[1], ".lal")
	if err != nil {
		fmt.Println("OMG")
		return
	}
	fmt.Printf("1 1 1 %+v\n", res)
	fmt.Println("-------------")
}