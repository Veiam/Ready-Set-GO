package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var pat = regexp.MustCompile(os.Args[3])

func main() {
	source, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	destination, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer destination.Close()

	scanner := bufio.NewScanner(source)
	writer := bufio.NewWriter(destination)

	for scanner.Scan() {
		fmt.Fprintf(writer, "%v\n", pat.FindString(scanner.Text()))
	}

	writer.Flush()
}
