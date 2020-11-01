package main

import (
	"fmt"
	jz "github.com/uakci/jvozba/v2"
	"os"
)

func jvozba(respond func(string), tanru string) {
	defer func() {
		if r := recover(); r != nil {
			respond("**spaji nabmi** .u’u")
			fmt.Fprint(os.Stderr, r)
		}
	}()

	lujvo, err := jz.Jvozba(tanru, jz.Brivla)
	if err != nil {
		respond("**nabmi**: " + err.Error())
	} else {
		respond(fmt.Sprintf("%s (%d)", lujvo, jz.Score(lujvo)))
	}
}
