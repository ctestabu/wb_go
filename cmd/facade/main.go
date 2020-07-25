//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("test123")
//}
//
////TODO: Простая логика с интерфейсами.
//
//type Converter interface {
//	//method Receive
//}
//
//type changer interface {
//	//change customFile
//	//list files
//}
//
//type customFile interface {
//	//giveConvertedFile
//	giveConvertedFile(fileName string) (err error)
//}
//
//type converter struct {
//	files customFile
//	toChange changer
//}
//
//func (c *converter) Receive(fileNames ...string) (msg string, err error) {
//	//logs?
//	//main control function
//	for _, names := range fileNames {
//		if err = c.files.giveConvertedFile(names);err !=nil {
//			return
//		}
//		//todo changefile(convert); list changet files; logs
//	}
//	return
//}
//
////Crestes new Converter entity
//func NewConverter(files customFile, toChange changer) Converter {
//	return &converter{
//		files: files,
//		toChange: toChange,
//	}
//}

//TODO: Порядок работы программы - Приходит строка с названием файла --- изменяется на нужный формат
//TODO --- заполняется список измененных файлов --- и создается файл с нужным название. Не забыть логи удалить потом файлы
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}