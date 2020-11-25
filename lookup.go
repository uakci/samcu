package samcu

import (
	"fmt"
	"strings"
)

const (
	lang      = "English"
	lowLimit  = 10
	highLimit = 50
)

var dict Dictionaries

type Dictionaries map[string]Dictionary

type Dictionary map[string]Definition

type Definition struct {
	Type       string   `json:"type"`
	Author     string   `json:"author"`
	Definition string   `json:"definition"`
	Notes      string   `json:"notes"`
	Glosses    []string `json:"glosses"`
}

func formatDef(a string, v Definition) string {
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

func sisku(cmd string, args []string) string {
	a := H.Replace(strings.Join(args, " "))
	if len(a) == 0 {
		return "need input"
	}

	dic := dict[lang]
	if len(cmd) > 6 {
		var ok bool
		dic, ok = dict[cmd[6:]]
		if !ok {
			return "no such dictionary as " + cmd[6:]
		}
	}

	vla, ok := dic[a]
	if ok {
		return formatDef(a, vla)
	}

	searched := "glosses"
	matches := map[string]string{}
	for head, vla := range dic {
		for i, gloss := range vla.Glosses {
			if gloss == a {
				parts := append(append(append([]string{}, vla.Glosses[:i]...), "**"+gloss+"**"), vla.Glosses[i+1:]...)
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
		searched = "definitions"
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
		searched = "definitions"
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
		return "facki tu’a no da"
	case len(matches) <= lowLimit:
		i := 0
		buil.WriteString("**in " + searched + "**")
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
		return "too many hits – try different query"
	}
	return buil.String()
}
