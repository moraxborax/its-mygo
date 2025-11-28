package main

import (
	"fmt"
	"log"
	"moraxborax/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"tomori", "penguin", "gugugaga"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	//message := greetings.Hello("Gladys")
	fmt.Println(messages)
}