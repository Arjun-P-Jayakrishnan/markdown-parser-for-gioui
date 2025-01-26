package handlers

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type FileHandler struct {
	FileName string
	FilePath string
	reader io.Reader
	//writer io.Writer
}


func (f FileHandler) openFile() (*os.File) {
	//join the files
	path := []string{f.FilePath,f.FileName};
	pathToFile := strings.Join(path,"")

	file,err :=os.Open(pathToFile)

	if err!=nil{
		log.Fatal(err)
		panic(err)
		
	}

	return file
}

/*
	@description reads data from file
*/
func (f FileHandler) ReadFromFile() (string, error) {
	buffer := make([]byte,1024)
	f.reader=f.openFile()


	line :=""
	lines := []string{}
	for {

		n,err := f.reader.Read(buffer)
		
		for  index,data := range buffer[:n] {
			if string(rune(data))=="\n" || index==n {
				fmt.Println(line)
				lines=append(lines,line)
			
				line=""
			}else {
				line+=string(rune(data))
			}
		}

		fmt.Println(lines)


		//check if reached the end of file
		if err==io.EOF{
			return "",nil
		}

		//if there is some other error
		if err!=nil {
			return "err",err
		}
	}
}

