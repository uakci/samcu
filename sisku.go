package samcu

import (
	"fmt"
	"github.com/uakci/samcu/jvs"
  "github.com/uakci/samcu/common"
	"strings"
	"sync"
)

var Sisku = Command{
  "sisku", sisku,
  "sisku fi tu'a lo'e vlaste",
  []CommandOption{
		{"selsisku", "valsi je cu poi'i do djica lo ka facki fi ke'a", nil, StringType},
  },
  []CommandOption{BanguOpt},
}

const (
	lowLimit  = 10
	highLimit = 75
)

var Index jvs.IndexType = jvs.IndexType{&sync.RWMutex{}, jvs.Dictionaries{}}

func GetDict(name string) (jvs.Dictionary, error) {
	Index.Mutex.RLock()
	defer Index.Mutex.RUnlock()
	d, o := Index.Index[name]
	if o {
		return d, nil
	} else {
		return nil, fmt.Errorf("fliba lo ka facki moi'a lo me zo'oi %s vlacku", name)
	}
}

func formatDef(a string, v jvs.Definition) string {
	notes := ""
	if v.Notes != "" {
		notes = fmt.Sprintf(" *%s*", v.Notes)
	}
	return fmt.Sprintf("**%s** [%s]: %s%s (%s)",
		a, v.Type, v.Definition, notes, v.Author)
}

func tryFind(where, what string) (string, bool) {
	if i := strings.Index(where, what); i != -1 {
		return where[:i] + "**" + what + "**" + where[i+len(what):], true
	} else {
		return "", false
	}
}

func sisku(args map[string]any) (string, error) {
  selsisku := args["selsisku"].(string)
  bangu := GetBangu(args)
	dic, err := GetDict(bangu)
	if err != nil {
		return "", err
	}

	a := common.ReplaceH(selsisku)
	vla, ok := dic[a]
	if ok {
		return formatDef(a, vla), nil
	}

	searched := "tordu velski"
	matches := map[string]string{}
	for head, vla := range dic {
		for i, gloss := range vla.Glosses {
			if gloss == a {
        parts := make([]string, 0, 3)
				parts = append(parts, vla.Glosses[:i]...)
        parts = append(parts, "**"+gloss+"**")
        parts = append(parts, vla.Glosses[i+1:]...)
				matches[head] = strings.Join(parts, ", ")
				break
			}
		}
		if len(matches) > highLimit {
			break
		}
	}

	if len(matches) == 0 {
		matches = map[string]string{}
		searched = "clani velski"
		for head, vla := range dic {
			add, ok := tryFind(vla.Definition, "__"+a+"__")
			if ok {
				matches[head] = add
				if len(matches) > highLimit {
					break
				}
			}
		}
	}

	if len(matches) == 0 {
		matches = map[string]string{}
		searched = "clani velski"
		for head, vla := range dic {
			add, ok := tryFind(vla.Notes, a)
			if ok {
				matches[head] = add
				if len(matches) > highLimit {
					break
				}
			}
		}
	}

	buil := strings.Builder{}
	switch {
	case len(matches) == 0:
		return "", fmt.Errorf("facki tu'a no da")
	case len(matches) <= lowLimit:
		i := 0
		buil.WriteString("**sisku fi lo'e " + searched + "**")
		for vla, match := range matches {
			buil.WriteString("\n")
			buil.WriteString(vla)
			buil.WriteString(": ")
			buil.WriteString(match)
			i++
		}
		if len(matches) == 1 {
			buil.WriteString("\n")
			for k := range matches {
				buil.WriteString(formatDef(k, dic[k]))
			}
		}
	case len(matches) <= highLimit:
		buil := strings.Builder{}
		i := 0
		for vla := range matches {
			if i > 0 {
				buil.WriteString(", ")
			}
			buil.WriteString(vla)
			i++
		}
	default:
		return "", fmt.Errorf("du'e da mapti lo jai se sisku pe do i ko troci lo ka cpedu su'o drata")
	}
	return buil.String(), nil
}
