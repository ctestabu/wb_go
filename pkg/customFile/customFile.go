package customFile

type CustomFile interface {
	//giveConvertedFile
	giveConvertedFile(fileName string) (err error)
}

type customFile struct {
	fileNames []string
}

func (c *customFile) giveConvertedFile(fileName string) (err error) {

}


func New(fileNames []string) CustomFile {
	return &customFile{
		fileNames: fileNames,
	}
}