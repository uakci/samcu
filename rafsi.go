package samcu

import (
	"fmt"
	"strings"

	"github.com/uakci/jvozba/v3"
	"github.com/uakci/samcu/common"
)

func marafsi(arg string) (string, error) {
	arg = common.ReplaceH(arg)

	if rafsi, ok := jvozba.Rafsi[arg]; ok {
		return fmt.Sprintf("ra'oi **%s** rafsi zo %s", strings.Join(rafsi, "** je ra'oi **"), arg), nil
	} else {
		return fmt.Sprintf("no da rafsi zo %s", arg), nil
	}
}

func rafsima(arg string) (string, error) {
	arg = common.ReplaceH(arg)

	var selrafsi string
	found := false
	for sr, rafsiporsi := range jvozba.Rafsi {
		for _, rafsi := range rafsiporsi {
			if rafsi == arg {
				selrafsi = sr
				found = true
				break
			}
		}
	}

	if found {
		return fmt.Sprintf("ra'oi %s rafsi zo **%s**", arg, selrafsi), nil
	} else {
		return fmt.Sprintf("ra'oi %s rafsi no da", arg), nil
	}
}
