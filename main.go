package main

import (
	"flag"
	"fmt"
)

func main() {

	filePath := flag.String("file", "", "Ruta del archivo de entrada")
	outputPath := flag.String("-o", "", "Ruta del archivo de salida")
	cleanFilePath := flag.String("clean", "", "Ruta del archivo de entrada para limpiar")
	flag.Parse()
	if *filePath != "" {
		Extractor(filePath, *outputPath)
	} else if *cleanFilePath != "" {
		RemoveDuplicates(*cleanFilePath)
	} else {
		fmt.Println("Por favor, proporciona la ruta del archivo usando la bandera -file o -clean para remover duplicados de un archivo")
		return
	}

}
