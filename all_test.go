package main

import (
	"regexp"
	"testing"

	"github.com/weidaicheng/github-action-sharing/service-greeting/svc"
)

func TestSayHello(t *testing.T) {
	// arrange
	name := "Jonathan"
	want := regexp.MustCompile(`\b` + name + `\b`)

	// act
	msg := svc.SayHello(name)

	// assert
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("%q") = %q, want match for %#q`, name, msg, want)
	}
}
