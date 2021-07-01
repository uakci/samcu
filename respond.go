package samcu

import (
	"strings"
)

var handlers = map[string]func(string, []string) string{
	"lujvo": lujvo,
	"rafsi": rafsi,
	"sisku": sisku,
	"katna": katna,
	"gloss": gloss,
	"parse": parse,

	"l": lujvo,
	"r": rafsi,
	"s": sisku,
	"k": katna,
	"g": gloss,
}

func Respond(data string) (string, bool) {
	fields := strings.Fields(data)
	if len(fields) == 0 {
		return "", false
	}
	cmd := strings.TrimSuffix(fields[0], ":")
	fields = fields[1:]

	handler, ok := handlers[cmd]
	if !ok {
		return "", false
	}
	return handler(cmd, fields), true
}
