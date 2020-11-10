package main

import (
	"fmt"
	"strings"
)

const (
	lang      = "English"
	lowLimit  = 10
	highLimit = 50
)

var dict dictionaries

type dictionaries map[string]dictionary

type dictionary map[string]definition

type definition struct {
	Type       string
	Author     string
	Definition string
	Notes      string
	Glosses    []string
}

func formatDef(a string, v definition) string {
	notes := ""
	if v.Notes != "" {
		notes = fmt.Sprintf(" *%s*", v.Notes)
	}
	return fmt.Sprintf("**%s** [%s]: %s%s (%s)",
		a, v.Type, v.Definition, notes, v.Author)
}

func lookup(respond func(string), _ string, args []string) {
	a := h.Replace(strings.Join(args, " "))
	vla, ok := dict[lang][a]
	if !ok {
		respond("facki tu’a no da")
	} else {
		respond(formatDef(a, vla))
	}
}

func tryFind(where, what string) (string, bool) {
	if i := strings.Index(where, what); i != -1 {
		return where[:i] + "**" + what + "**" + where[i+len(what):], true
	} else {
		return "", false
	}
}

func reverseLookup(respond func(string), _ string, args []string) {
	a := h.Replace(strings.Join(args, " "))
	if len(a) == 0 {
		respond("need input")
		return
	}
	searched := "glosses"
	matches := map[string]string{}
	for head, vla := range dict[lang] {
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
		for head, vla := range dict[lang] {
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
		for head, vla := range dict[lang] {
			add, ok := tryFind(vla.Notes, a)
			if ok {
				matches[head] = add
				if len(matches) > highLimit {
					break
				}
			}
		}
	}
	switch {
	case len(matches) == 0:
		respond("facki tu’a no da")
	case len(matches) <= lowLimit:
		buil := strings.Builder{}
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
				buil.WriteString(formatDef(k, dict[lang][k]))
			}
		}
		respond(buil.String())
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
		respond(buil.String())
	default:
		respond("too many hits – try different query")
	}
}
