package samcu

import (
	"fmt"
  "log"
	"github.com/uakci/jvozba/v3"
	"github.com/uakci/samcu/common"
	"strings"
)

var Rafsi = Command{
  "rafsi", rafsi,
	"tutci je srana be lo si'o rafsi",
  []CommandOption{
    {"fiha", "do xa'o djuno fi lo'e rafsi ji se rafsi",
    [][2]string{
      {"fa", "do cusku lo rafsi i mi cusku lo se rafsi"},
      {"fe", "do cusku lo se rafsi i mi cusku lo rafsi"},
    }, StringType},
		{"zo", "valsi je poi'i do xa'o djuno", nil, StringType},
  },
  []CommandOption{},
}

func rafsi(args map[string]any) (string, error) {
  fiha := args["fiha"].(string)
  ma := args["zo"].(string)
	ma = common.ReplaceH(ma)

	switch fiha {
	case "fe":
		builder := strings.Builder{}
		if r, ok := jvozba.Rafsi[ma]; ok {
			for i, rafsi := range r {
				if i != 0 {
					builder.WriteString(" je ")
				}
				fmt.Fprintf(&builder, "ra'oi **%s**", rafsi)
			}
		} else {
			builder.WriteString("no da")
		}
		fmt.Fprintf(&builder, " rafsi zo %s", ma)
		return builder.String(), nil

	case "fa":
		var selrafsi *string
		for sr, rafsiporsi := range jvozba.Rafsi {
			for _, rafsi := range rafsiporsi {
				if rafsi == ma {
					selrafsi = &sr
					break
				}
			}
		}

		if selrafsi != nil {
			return fmt.Sprintf("zo **%s** se rafsi ra'oi %s", *selrafsi, ma), nil
		} else {
			return fmt.Sprintf("no da se rafsi ra'oi %s", ma), nil
		}
	}

  log.Panicf("unknown fiha %s", fiha)
  return "", nil
}
