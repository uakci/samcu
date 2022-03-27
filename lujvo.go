package samcu

import (
	"strings"

	"github.com/uakci/jvozba/v3"
)

func lujvo(args []string) (string, error) {
	return jvozba.Jvozba(strings.Join(args, " "), jvozba.Brivla)
}
