package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/jasonpaulos/tealx/language/compiler"
	"github.com/jasonpaulos/tealx/language/element"
)

func parseProgram(file *os.File) (*element.Program, error) {
	e, err := element.UnmarshalXml(file)
	if err != nil {
		return nil, err
	}
	if program, ok := e.(*element.Program); ok {
		return program, nil
	}
	return nil, fmt.Errorf("expected type Program but got %#T", e)
}

func main() {
	parser := argparse.NewParser("tealx", "Compiles Tealx programs to TEAL")
	inFile := parser.File("i", "in", os.O_RDONLY, 0666, &argparse.Options{Required: true, Help: "Input file"})
	outFile := parser.File("o", "out", os.O_WRONLY|os.O_CREATE, 0666, &argparse.Options{Required: true, Help: "Output file"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	program, err := parseProgram(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when parsing %s: %v\n", inFile.Name(), err)
		os.Exit(1)
	}

	err = compiler.Compile(*program, outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when compiling: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled %s to %s\n", inFile.Name(), outFile.Name())
}
