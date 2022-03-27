package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/uakci/samcu"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func comment(ok bool, msg string, err error, invocation []string) {
	if !ok {
		fmt.Printf("la'oi %s na slabu mi\n", invocation[0])
	} else if err != nil {
		fmt.Printf("la'e di'e cfila: %s\n", err.Error())
	} else {
		fmt.Println(msg)
	}
}

func main() {
	samcu.Init()

	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in := strings.Fields(scanner.Text())
			if len(in) == 0 {
				break
			}
			ok, msg, err := samcu.Respond(in)
			comment(ok, msg, err, in)
		}
	} else {
		ok, msg, err := samcu.Respond(os.Args[1:])
		comment(ok, msg, err, os.Args[1:])
		switch {
		case !ok:
			os.Exit(2)
		case err != nil:
			os.Exit(1)
		}
	}
}
