package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name Is Empty")
	}

	message := fmt.Sprintf("Hi, %s. Welcome!", name)
	return message, nil
}
