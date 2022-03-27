package samcu

import (
	"fmt"
	"strings"

	"github.com/uakci/samcu/jvs"
)

func NoEmpty(inner Handler) Handler {
	return func(args []string) (string, error) {
		if len(args) == 0 {
			return "", fmt.Errorf("ko cusku su'o datni")
		}
		return inner(args)
	}
}

func Repeatedly(inner func(string) (string, error)) Handler {
	return NoEmpty(func(args []string) (string, error) {
		switch len(args) {
		case 1:
			return inner(args[0])
		default:
			var builder strings.Builder
			for _, arg := range args {
				if res, err := inner(arg); err != nil {
					fmt.Fprintf(&builder, "%s: ⚠ %s\n", arg, err.Error())
				} else {
					fmt.Fprintf(&builder, "%s: %s\n", arg, res)
				}
			}
			return builder.String(), nil
		}
	})
}

func WithDictionary(inner func(jvs.Dictionary) Handler) Handler {
	return func(args []string) (string, error) {
		bangu := defaultBangu
		if len(args) > 0 && strings.HasPrefix(args[0], "-") {
			bangu = args[0][1:]
			args = args[1:]
		}

		Index.Mutex.RLock()
		dict, ok := Index.Index[bangu]
		if ok {
			Index.Mutex.RUnlock()
			return inner(dict)(args)
		} else {
			vlacku := make([]string, 0, len(Index.Index))
			for k := range Index.Index {
				vlacku = append(vlacku, k)
			}
			Index.Mutex.RUnlock()
			return "", fmt.Errorf(
				"la'oi %s na vlacku .i ko cuxna fi la'e di'e fa'o: %s",
				bangu, strings.Join(vlacku, ", "))
		}
	}
}
