package main

import (
	jz "github.com/uakci/jvozba/v2"
	"strings"
)

func rafsi(respond func(string), cmd string, args []string) {
	if len(args) != 1 {
		respond("one argument expected")
		return
	}
	arg := h.Replace(args[0])
	var results []string
	if cmd == "selrafsi" {
		for selrafsi, rafsiporsi := range jz.Rafsi {
			for _, rafsi := range rafsiporsi {
				if rafsi == arg {
					results = append(results, selrafsi)
				}
			}
		}
	} else {
		results, _ = jz.Rafsi[arg]
	}
	if len(results) == 0 {
		respond("no da")
	} else {
		respond(strings.Join(results, ", "))
	}
}
