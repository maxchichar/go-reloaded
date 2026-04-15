package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"go-reloaded/functions"
)

func main()  {
	if len(os.Args) != 3 {
		fmt.Println("✗ Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if (!(strings.HasSuffix(inputFile, ".txt"))) || (!(strings.HasSuffix(outputFile, ".txt"))){
		fmt.Println("✗ File name has to end with .txt")
		os.Exit(1)
	}

	if inputFile == outputFile{
		fmt.Println("✗ File name cannot be the same")
		os.Exit(1)
	}

	file, err := os.OpenFile(inputFile, os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal("✗ Error opening file:",err)
	}
	defer file.Close()

	var b strings.Builder


	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
	}


	if err := scanner.Err() ;err != nil {
		log.Fatalln("✗ Error reading file:", err)
	}

	fullContent := b.String()

	// === USE FUNCTIONS FROM PACKAGE ===
	tokens := functions.SplitElement(fullContent)
	tokens = functions.ProcessCase(tokens)
	tokens = functions.ProcessQuotes(tokens)
	tokens = functions.ProcessHexBin(tokens)
	tokens = functions.ProcessPunctuation(tokens)
	tokens = functions.ProcessAnRule(tokens)

	finalContent := strings.Join(tokens, " ")
	//<=======================================>
	
	file1, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	_, err = file1.WriteString(finalContent)
	if err != nil {
		log.Fatalf("✗ Error writing output: %v", err)
	}
	
	fmt.Println("Reloading completed...")
}