package handlers

import (
	"log"
	"os"
	"bufio"
)

/*
	File Handler structure creates a handler abstraction for handling file operations 
*/
type FileHandler struct {
	FileName string // Name of the file
	FilePath string // Path to the given file
}





/*
	Reads data from a file and returns the contents of the file or error
*/
func (f FileHandler) ReadFile() ([]string, error) {
	file,f_err :=os.Open(f.FilePath+f.FileName)

	if f_err!=nil {
		log.Fatal(f_err)
		panic(f_err)
		
	}

	defer file.Close()

	lines :=make([]string,0)

	//Initialize a buffer for holding data
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line:=scanner.Text()
		lines=append(lines,line)
		
	}


	return lines,nil
}


