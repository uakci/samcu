package samcu

import (
	"fmt"
	jvozba "github.com/uakci/jvozba/v3"
	"os"
	"strings"
)

func lujvo(_ string, tanru []string) (resp string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, r)
			resp = "**spaji nabmi** .u’u"
		}
	}()

	jvo, err := jvozba.Jvozba(strings.Join(tanru, " "), jvozba.Brivla)
	if err != nil {
		resp = fmt.Sprintf("**nabmi**: %s", err.Error())
	} else {
    resp = jvo
  }
	return
}
