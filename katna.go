package main

import (
	jz "github.com/uakci/jvozba/v3"
	"strings"
)

func katna(respond func(string), _ string, args []string) {
	if len(args) != 1 {
		respond("one argument expected")
		return
	}
	decomp, e := jz.Veljvo(args[0])
	if e != nil {
		respond(e.Error())
	} else {
		respond(strings.Join(decomp, " "))
	}
}
