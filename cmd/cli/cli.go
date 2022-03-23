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

func main() {
	samcu.Init()

	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in := scanner.Text()
			if len(in) == 0 {
				break
			}
			response, ok := samcu.Respond(in)
			if !ok {
				fmt.Printf("unknown command %s\n", strings.SplitN(in, " ", 1)[0])
			} else {
				fmt.Println(response)
			}
		}
	} else {
		response, ok := samcu.Respond(strings.Join(os.Args[1:], " "))
		if !ok {
			fmt.Printf("unknown command %s\n", os.Args[1])
			os.Exit(1)
		}
		fmt.Println(response)
	}
}
