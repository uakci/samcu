package samcu

import (
	"fmt"
	jvozba "github.com/uakci/jvozba/v3"
	"strings"
)

func rafsi(cmd string, args []string) string {
	if len(args) != 1 {
		return "one argument expected"
	}
	arg := H.Replace(args[0])
  var bits []string

	if r, ok := jvozba.Rafsi[arg]; ok {
		bits = append(bits, fmt.Sprintf("%s → {%s}", arg, strings.Join(r, ", ")))
	}

	var (selrafsi string; selselrafsi []string)
	for sr, rafsiporsi := range jvozba.Rafsi {
		for _, rafsi := range rafsiporsi {
			if rafsi == arg {
				selrafsi = sr
        selselrafsi = rafsiporsi
				break
			}
		}
	}

	if len(selselrafsi) > 0 {
		bits = append(bits, fmt.Sprintf("%s → {%s}", selrafsi, strings.Join(selselrafsi, ", ")))
	}

  if len(bits) > 0 {
    if len(bits) == 2 && bits[0] == bits[1] {
      return bits[0]
    } else {
      return strings.Join(bits, "; ")
    }
  } else {
    return "∅"
  }
}
