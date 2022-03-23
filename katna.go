package samcu

import (
	jvozba "github.com/uakci/jvozba/v3"
	"strings"
)

var Katna = Command{
  "katna", katna,
	"lanli lo du'u ma kau pagbu lo'e lujvo",
  []CommandOption{
    {"lujvo", "se katna pe'a", nil, StringType},
  },
  []CommandOption{},
}

func katna(args map[string]any) (string, error) {
  lujvo := args["lujvo"].(string)
	resp := strings.Builder{}
	for _, v := range strings.Split(lujvo, " ") {
		decomp, e := jvozba.Veljvo(v)
		if e != nil {
			return "", e
		}
    resp.WriteString(v)
    resp.WriteString(" → ")
		resp.WriteString(strings.Join(decomp, " "))
		resp.WriteByte('\n')
	}
	return resp.String(), nil
}
