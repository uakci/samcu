package samcu

import (
	jvozba "github.com/uakci/jvozba/v3"
	"strings"
)

func katna(_ string, args []string) string {
	if len(args) != 1 {
		return "one argument expected"
	}
	decomp, e := jvozba.Veljvo(args[0])
	if e != nil {
		return e.Error()
	}
	return strings.Join(decomp, " ")
}
