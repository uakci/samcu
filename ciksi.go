package samcu

import (
	"strings"

	"github.com/uakci/jvozba/v3"
	"github.com/uakci/samcu/common"
	"github.com/uakci/samcu/jvs"
)

func ciksi(dict jvs.Dictionary) Handler {
	return NoEmpty(func(args []string) (string, error) {
		return ciksi_(dict, args)
	})
}

func ciksi_(dict jvs.Dictionary, args []string) (string, error) {
	var result strings.Builder
	for i, a := range args {
		if i > 0 {
			result.WriteRune('\u2003')
		}
		a = common.ReplaceH(a)

		var subject []string
		if jvozba.IsGismu([]byte(a)) || jvozba.IsCmavo([]byte(a)) || len(jvozba.Katna([]byte(a))) == 1 {
			subject = []string{a}
		} else {
			var e error
			subject, e = jvozba.Veljvo(a)
			if e != nil {
				subject = []string{a}
			}
		}

		for j, s := range subject {
			if j > 0 {
				result.WriteRune('—')
			}
			def, ok := dict[s]
			if !ok || len(def.Glosses) == 0 {
				result.WriteString("*" + s + "*")
			} else {
				result.WriteString(def.Glosses[0])
			}
		}
	}
	return result.String(), nil
}
