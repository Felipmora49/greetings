package greetings

import (

	"testing"
	"regexp"
)

func TestHelloName(t *testing.T){

	name := "Julian"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Julian")

	if !want.MatchString(msg) || err != nil {

		t.Fatalf(`Hello("Julian") = %q, %v, quiere coincidencia para %#q, nil`, msg , err , want)
	}
}

func TestHelloEmpty(t * testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg , err )
	}
}