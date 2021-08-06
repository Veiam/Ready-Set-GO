package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefiend Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())

	// // Get a greeting message and print it.
	// message, nil := greetings.Hello("Stranger")

	// A slice of names.
	names := []string{"Partner", "Pard", "Part"}

	// Request a greeting message.
	message, err := greetings.Hellos(names)
	// If an error was returend, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returend, print the returned message
	// to the console.
	fmt.Println(message)
}
