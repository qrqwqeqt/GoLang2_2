package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/qrqwqeqt/GoLang2_2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Path to an input file")
	outputFile      = flag.String("o", "", "Path to an output file")
)

func getReadFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getWriteFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	flag.Parse()
	if *inputExpression != "" && *inputFile != "" {
		log.Fatal("Multiple input flags specified")
	}

	var (
		reader io.Reader
		writer io.Writer
	)

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file := getReadFile(*inputFile)
		defer file.Close()
		reader = file
	} else {
		log.Fatal("No input flags specified")
	}

	if *outputFile != "" {
		file := getWriteFile(*outputFile)
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err := handler.Compute()
	if err != nil {
		log.Fatal(err) //Outputs to stderr
	}
}
