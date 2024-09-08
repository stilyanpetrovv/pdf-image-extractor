package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// Prompt the user for the PDF file path
	fmt.Print("Enter the name of the PDF file to extract images from: ")

	// Read the file path from user input
	reader := bufio.NewReader(os.Stdin)
	pdfPath, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Trim any extra newline characters from the input
	pdfPath = pdfPath[:len(pdfPath)-1]

	// Check if the file exists
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		log.Fatalf("Error: The file '%s' does not exist.", pdfPath)
	}

	// Directory where the extracted images will be saved
	outputDir := "output_images"

	// Create the output directory if it doesn't exist
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	// Extract images from the PDF and save them in the output directory
	err = api.ExtractImagesFile(pdfPath, outputDir, nil, nil)
	if err != nil {
		log.Fatalf("Error extracting images from PDF: %v", err)
	}

	fmt.Println("Images extracted and saved successfully!")
}

