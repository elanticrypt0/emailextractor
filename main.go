package main

import (
	"flag"
	"fmt"
)

func main() {

	filePath := flag.String("file", "", "Path of the input file")
	outputPath := flag.String("o", "", "Path of the output files")
	maxBuffer := flag.Int("buffer", 200, "Buffer's size in MB")
	cleanFilePath := flag.String("clean", "", "Path of the file to remove duplicates")
	flag.Parse()
	if *filePath != "" {
		Extractor(filePath, *outputPath, *maxBuffer)
	} else if *cleanFilePath != "" {
		RemoveDuplicates(*cleanFilePath)
	} else {
		fmt.Println("Please, write the file's path using the -file flag or -clean to remove duplicates in a file.")
		return
	}

}
