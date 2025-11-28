package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	//broken code
	//message := fmt.Sprint(randomFormat())
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
	
}


//modifying old function might break old code so lets write a new stuff

func Hellos (names []string) (map[string]string, error) {
	messages := make(map[string]string)
	//make is the constructor?
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
		// map is basically python's dict in golang
		
	}
	return messages, nil
}