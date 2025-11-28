package main

import (
	"fmt"

	"moraxborax/greetings"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}