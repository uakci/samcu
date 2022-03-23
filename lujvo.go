package samcu

import (
	"fmt"
	"github.com/uakci/jvozba/v3"
	"os"
)

var Lujvo = Command{
  "lujvo", lujvo,
	"kanji lo tarmi xagrai lujvo lo'e ve lujvo",
  []CommandOption{
    {"kratau", "du'u lujvo fo ma kau", nil, StringType},
  },
  []CommandOption{},
}

func lujvo(args map[string]any) (resp string, err error) {
  tanru := args["kratau"].(string)
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
			err = fmt.Errorf("**spaji nabmi** .u'u")
		}
	}()
	resp, err = jvozba.Jvozba(tanru, jvozba.Brivla)
  if err == nil {
    resp = fmt.Sprintf("%s → %s", tanru, resp)
  }
	return
}
