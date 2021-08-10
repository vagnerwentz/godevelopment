package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

// TestHelloName calls greetings. Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	fmt.Println(want)
	message, err := Hello("Gladys")
	fmt.Println(!want.MatchString(message))
	fmt.Println(want.MatchString(message))
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys")') = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}