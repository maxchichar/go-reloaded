package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main()  {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile{
		fmt.Println("✗ Files cannot be the same")
		os.Exit(1)
	}

	file, err := os.OpenFile(inputFile, os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var b strings.Builder


	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
	}


	if err := scanner.Err() ;err != nil {
		log.Fatalln("Reading standard input:", err)
	}

	
	file1, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	_, err = file1.WriteString(b.String())
	if err != nil {
		log.Fatalf("Error writing content: %v", err)
	}
	
}