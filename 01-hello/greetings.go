package greetings

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	message := fmt.Sprintf("%v, %v. Welcome!", englishHelloPrefix, name)
	return message
}
