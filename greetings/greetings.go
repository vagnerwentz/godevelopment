package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Returns a greeting that embeds the name in a message.
	// The same about var message string - message = ...
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	//fmt.Println(message)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		//fmt.Println(index)
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
		//fmt.Println(messages[name])

	}
	return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
// randomFormat starts with a lowercase letter, making it accessible only to code in its own package
func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}