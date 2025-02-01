package main

import "github.com/Arjun-P-Jayakrishnan/markdown-parser-for-gioui/handlers"

func main() {

	fileHandler:=handlers.FileHandler{
        FileName: "mark.md",
        FilePath: "",
    } 

    mdParser := handlers.Parser{}

    mdParser.ConvertFileToMarkdown(fileHandler)
}