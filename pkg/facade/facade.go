package facade

//local
type changer interface {
	//change customFile
	//list files
}

//local
type file interface {
	//giveConvertedFile
	giveConvertedFile(fileName string) (err error)
}

type Converter interface {
	//method Receive
	Receive(fileNames ...string) (msg string, err error)
}

type converter struct {
	files    file
	toChange changer
}

func (c *converter) Receive(fileNames ...string) (msg string, err error) {
	//logs?
	//main control function
	for _, names := range fileNames {
		if err = c.files.giveConvertedFile(names); err != nil {
			return
		}
		//todo changefile(convert); list changet files; logs
	}
	return
}

//Crestes new Converter entity
func NewConverter(files file, toChange changer) Converter {
	return &converter{
		files:    files,
		toChange: toChange,
	}
}
