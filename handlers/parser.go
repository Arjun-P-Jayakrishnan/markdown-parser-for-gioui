package handlers

import (
	"fmt"
	"log"
)

type Parser struct {
}

func (p Parser) ConvertFileToMarkdown(f FileHandler) error {

	lines, err := f.ReadFile()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
	return nil
}