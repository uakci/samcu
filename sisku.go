package samcu

import (
	"fmt"
	"strings"
	"sync"

	"github.com/uakci/samcu/common"
	"github.com/uakci/samcu/jvs"
)

const (
	inliningLimit = 10
)

var Index jvs.IndexType = jvs.IndexType{
	Mutex: &sync.RWMutex{},
	Index: jvs.Dictionaries{},
}

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

func sisku(dict jvs.Dictionary) Handler {
	return NoEmpty(func(args []string) (string, error) {
		return sisku_(dict, common.ReplaceH(strings.Join(args, " ")))
	})
}

func sisku_(dict jvs.Dictionary, term string) (string, error) {
	if vla, ok := dict[term]; ok {
		return formatDef(term, vla), nil
	}

	matches := map[string]string{}
	for head, vla := range dict {
		for i, gloss := range vla.Glosses {
			if gloss == term {
				parts := make([]string, 0, 3)
				parts = append(parts, vla.Glosses[:i]...)
				parts = append(parts, fmt.Sprintf("**%s**", gloss))
				parts = append(parts, vla.Glosses[i+1:]...)
				matches[head] = strings.Join(parts, ", ")
				break
			}
		}
	}

	if len(matches) == 0 {
		for head, vla := range dict {
			if add, ok := tryFind(vla.Notes, fmt.Sprintf("__%s__", term)); ok {
				matches[head] = add
			}
		}
	}

	if len(matches) == 0 {
		for head, vla := range dict {
			if add, ok := tryFind(vla.Definition, term); ok {
				matches[head] = add
			} else if add, ok = tryFind(vla.Notes, term); ok {
				matches[head] = add
			}
		}
	}

	switch {
	case len(matches) == 0:
		return "", fmt.Errorf("facki tu'a no da")

	case len(matches) <= inliningLimit:
		var builder strings.Builder
		for vla, match := range matches {
			fmt.Fprintf(&builder, "%s: %s\n", vla, match)
			fmt.Fprintln(&builder, formatDef(vla, dict[vla]))
		}
		return builder.String(), nil

	default:
		headwords := make([]string, 0, len(matches))
		for vla := range matches {
			headwords = append(headwords, vla)
		}
		return strings.Join(headwords, ", "), nil
	}
}
