package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/uakci/samcu"
	"os"
	"regexp"
	"strings"
)

type xmlJbovlaste struct {
	Directions []direction `xml:"direction"`
}

type direction struct {
	From    string  `xml:"from,attr"`
	To      string  `xml:"to,attr"`
	Valsi   []valsi `xml:"valsi"`
	Lemmata []lemma `xml:"nlword"`
}

type valsi struct {
	Word         string      `xml:"word,attr"`
	Type         string      `xml:"type,attr"`
	Selmaho      *string     `xml:"selmaho"`
	User         user        `xml:"user"`
	Definition   string      `xml:"definition"`
	DefinitionID int         `xml:"definitionid"`
	Notes        *string     `xml:"notes"`
	Glosses      []glossWord `xml:"glossword"`
}

type user struct {
	UserName string `xml:"username"`
	RealName string `xml:"realname"`
}

type glossWord struct {
	Word  string `xml:"word,attr"`
	Sense string `xml:"sense,attr"`
}

type lemma struct {
	Word  string `xml:"word,attr"`
	Sense string `xml:"sense,attr"`
	Valsi string `xml:"valsi,attr"`
}

func main() {
	grand := samcu.Dictionaries{}
	for _, fname := range os.Args[1:] {
		f, e := os.Open(fname)
		if e != nil {
			panic(e)
		}
		var jvs xmlJbovlaste
		e = xml.NewDecoder(f).Decode(&jvs)
		if e != nil {
			panic(e)
		}

		this := parseJVS(jvs)
		for k, v := range this {
			grand[k] = v
		}
	}
	f, e := os.Create("jvs.json")
	if e != nil {
		panic(e)
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	e = enc.Encode(grand)
	if e != nil {
		panic(e)
	}
}

var (
	linkRegex              = regexp.MustCompile(`\{(.+?)\}`)
	specialCharactersRegex = regexp.MustCompile("[\\`*_<>!|]")
)

func parseDollars(s string, embellishItalics bool) string {
	result := strings.Builder{}
	result.Grow(len(s))
	splinters := strings.Split(s, "$")
	for inter, splinter := range splinters {
		if inter%2 == 0 {
			splinter = specialCharactersRegex.ReplaceAllString(splinter, `\${0}`)
			splinter = linkRegex.ReplaceAllString(splinter, `__${1}__`)
			result.WriteString(splinter)
			continue
		}
		buf, succ, afterUnderscore := strings.Builder{}, true, false
		buf.Grow(len(splinter))
		for _, c := range splinter {
			switch {
			case (c >= 'a' && c <= 'z') || c == '\'':
				buf.WriteByte(byte(c))
				afterUnderscore = false
			case c == '=':
				buf.WriteString("\u2009=\u2009")
				afterUnderscore = false
			case c == '_':
				afterUnderscore = true
			case c == '{':
				if !afterUnderscore {
					succ = false
				}
			case c == '}':
				if afterUnderscore {
					afterUnderscore = false
				} else {
					succ = false
				}
			case c >= '0' && c <= '9':
				if afterUnderscore {
					buf.WriteRune(c - 0x30 + 0x2080)
				} else {
					buf.WriteByte(byte(c))
				}
			case c == ' ':
				afterUnderscore = false
			default:
				succ = false
			}
			if !succ {
				break
			}
		}
		if succ {
			if embellishItalics {
				result.WriteByte('*')
			}
			result.WriteString(buf.String())
			if embellishItalics {
				result.WriteByte('*')
			}
		} else {
			result.WriteByte('`')
			result.WriteString(splinter)
			result.WriteByte('`')
		}
	}
	return result.String()
}

func parseJVS(jvs xmlJbovlaste) samcu.Dictionaries {
	result := map[string]samcu.Dictionary{}
	for _, direction := range jvs.Directions {
		if len(direction.Lemmata) > 0 {
			continue
		}
		key := direction.To
		dict := map[string]samcu.Definition{}
		for _, valsi := range direction.Valsi {
			var typ string
			if valsi.Selmaho != nil {
				typ = *valsi.Selmaho
			} else {
				typ = valsi.Type
			}
			var notes string
			if valsi.Notes != nil {
				notes = *valsi.Notes
			} else {
				notes = ""
			}
			glosses := []string{}
			for _, g := range valsi.Glosses {
				if g.Sense != "" {
					glosses = append(glosses, fmt.Sprintf("%s (%s)", g.Word, g.Sense))
				} else {
					glosses = append(glosses, g.Word)
				}
			}
			dict[samcu.H.Replace(valsi.Word)] = samcu.Definition{
				Type:       typ,
				Author:     valsi.User.UserName,
				Definition: parseDollars(valsi.Definition, true),
				Notes:      parseDollars(notes, false),
				Glosses:    glosses,
			}
		}
		result[key] = dict
	}
	return result
}
