package samcu

import (
	"github.com/uakci/jvozba/v3"
  "github.com/uakci/samcu/common"
	"strings"
)

var Ciksi = Command{
  "ciksi", ciksi,
  "ciksi lo jufra fo tu'a su'o na lojbo",
  []CommandOption{
    {"jufra", "poi'i do djica lo nu mi ciksi ke'a", nil, StringType},
  },
  []CommandOption{BanguOpt},
}

func ciksi(args map[string]any) (string, error) {
  jufra := args["jufra"].(string)
  bangu := GetBangu(args)
  dict, err := GetDict(bangu)
	if err != nil {
		return "", err
	}

	result := strings.Builder{}
	for i, a := range strings.Fields(jufra) {
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
