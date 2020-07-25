package facade

//local
type changer interface {
	ChangeFile(fileName, format string) (msg string, err error)
	ListFilesToChange() (msg string, err error)
}

//local
type customFile interface {
	//giveConvertedFile
	GiveConvertedFile(fileName string) (err error)
}

type Converter interface {
	//method Receive
	Receive(fileNames ...string) (msg string, err error)
}

type converter struct {
	files    customFile
	toChange changer
}

func (c *converter) Receive(fileNames ...string) (msg string, err error) {
	//logs?
	//main control function
	for _, names := range fileNames {
		if err = c.files.GiveConvertedFile(names); err != nil {
			return
		}

	}
	return
}

//Creates new Converter entity
func NewConverter(files customFile, toChange changer) Converter {
	return &converter{
		files:    files,
		toChange: toChange,
	}
}
