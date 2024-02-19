package main

import (
	"regexp"
	"testing"

	"github.com/weidaicheng/service-greeting/svc"
)

func TestSayHello(t *testing.T) {
	// arrange
	name := "name"
	want := regexp.MustCompile(`\b` + name + `\b`)

	// act
	msg := svc.SayHello(name)

	// assert
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("%q") = %q, want match for %#q`, name, msg, want)
	}
}
