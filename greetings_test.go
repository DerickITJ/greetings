package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T){
	name:="Dalet"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg,err := Hello("Dalet")
	if !want.MatchString(msg) || err != nil{
		t.Fatalf(`Hello("Dalet") = %q,%v quiere coincidencia para %#q,nill`,msg,err,want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg,err := Hello("")
	if msg != "" || err == nil{
		t.Fatalf(`Hello("") = %q,%v quiere coincidencia para "",error`,msg,err)
	}
}