package main

import (
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

func main() {
	must(samcu.LoadJVS())

	if len(os.Args) < 2 {
		fmt.Println("need argument")
		os.Exit(1)
	}

  response, ok := samcu.Respond(strings.Join(os.Args[1:], " "))
	if !ok {
		fmt.Printf("unknown command %s\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println(response)
}
