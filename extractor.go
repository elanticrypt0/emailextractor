package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Extractor(filePath *string, outputPath string, maxBuffer int) {

	// Verificar si se proporcionó la ruta del archivo
	if *filePath == "" {
		fmt.Println("Por favor, proporciona la ruta del archivo usando la bandera -file")
		return
	}

	allEmailsFile := "all_emails.csv"
	domainsFile := "email_domains.csv"

	if outputPath != "" {
		allEmailsFile = outputPath + "/" + allEmailsFile
		domainsFile = outputPath + "/" + domainsFile
	}

	// Abrir el archivo de entrada
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Crear expresión regular para encontrar emails
	emailRegex := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)

	// Crear slice para almacenar todos los emails
	allEmails := []string{}

	// Crear slice para almacenar dominios inválidos
	emailDomainsSlice := []string{}

	// Crear scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Calcular el tamaño máximo del buffer en bytes (300 MB)
	maxBufferSize := maxBuffer * 1024 * 1024

	// Configurar el tamaño máximo del buffer del scanner
	scanner.Buffer(make([]byte, maxBufferSize), maxBufferSize)

	// Leer el archivo línea por línea
	for scanner.Scan() {
		line := scanner.Text()
		emails := emailRegex.FindAllString(line, -1)

		// Iterar sobre los emails encontrados en la línea
		for _, email := range emails {
			allEmails = append(allEmails, strings.ToLower(email))

			// Extraer el dominio del email
			parts := strings.Split(email, "@")
			if len(parts) == 2 && parts[1] != "gmail.com" && parts[1] != "hotmail.com" && parts[1] != "yahoo.com" {
				emailDomainsSlice = append(emailDomainsSlice, strings.ToLower(parts[1]))
			}
		}
	}

	// Verificar errores de escaneo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo línea por línea: %v\n", err)
		return
	}

	// Crear el archivo CSV de todos los emails

	allEmailsCSVFile, err := os.Create(allEmailsFile)
	if err != nil {
		fmt.Printf("Error al crear el archivo de resultados: %v\n", err)
		return
	}
	defer allEmailsCSVFile.Close()

	// Escribir todos los emails en el archivo CSV
	allEmailsWriter := csv.NewWriter(allEmailsCSVFile)
	defer allEmailsWriter.Flush()

	for _, email := range allEmails {
		if err := allEmailsWriter.Write([]string{email}); err != nil {
			fmt.Printf("Error al escribir en el archivo CSV: %v\n", err)
			return
		}
	}

	fmt.Println("Se han guardado todos los emails en el archivo all_emails.csv")

	// Crear el archivo CSV de dominios inválidos
	emailsDomainsCSVFile, err := os.Create(domainsFile)
	if err != nil {
		fmt.Printf("Error al crear el archivo de resultados: %v\n", err)
		return
	}
	defer emailsDomainsCSVFile.Close()

	// Escribir los dominios inválidos en el archivo CSV
	badDomainsWriter := csv.NewWriter(emailsDomainsCSVFile)
	defer badDomainsWriter.Flush()

	for _, domain := range emailDomainsSlice {
		if err := badDomainsWriter.Write([]string{domain}); err != nil {
			fmt.Printf("Error al escribir en el archivo CSV: %v\n", err)
			return
		}
	}

	fmt.Println("Se han guardado los dominios inválidos en el archivo ", domainsFile)
}
