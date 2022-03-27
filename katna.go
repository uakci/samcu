package samcu

import (
	"strings"

	"github.com/uakci/jvozba/v3"
)

func katna(arg string) (string, error) {
	if decomp, e := jvozba.Veljvo(arg); e != nil {
		return "", nil
	} else {
		return strings.Join(decomp, " "), nil
	}
}
