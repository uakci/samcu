package samcu

import (
	jvozba "github.com/uakci/jvozba/v3"
	"strings"
)

func gloss(_ string, args []string) string {
	result := strings.Builder{}
	for i, a := range args {
		if i > 0 {
			result.WriteRune('\u2003')
		}
		a = H.Replace(a)
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
			def, ok := dict[lang][s]
			if !ok || len(def.Glosses) == 0 {
				result.WriteString("*" + s + "*")
			} else {
				result.WriteString(def.Glosses[0])
			}
		}
	}
	return result.String()
}
