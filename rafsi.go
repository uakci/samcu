package main

import (
	"bytes"
	jz "github.com/uakci/jvozba/v2"
	"strings"
)

func rafsi(respond func(string), cmd string, args []string) {
	if len(args) != 1 {
		respond("one argument expected")
		return
	}
	arg := h.Replace(args[0])
	var results []string
	if cmd == "selrafsi" {
		for selrafsi, rafsiporsi := range jz.Rafsi {
			for _, rafsi := range rafsiporsi {
				if rafsi == arg {
					results = append(results, selrafsi)
				}
			}
		}
	} else {
		results, _ = jz.Rafsi[arg]
	}
	if len(results) == 0 {
		respond("no da")
	} else {
		respond(strings.Join(results, ", "))
	}
}

func katna(respond func(string), _ string, args []string) {
	if len(args) != 1 {
		respond("one argument expected")
		return
	}
	arg := h.Replace(args[0])
	rafpoi := jz.Katna([]byte(arg))
	tanru := make([]string, len(rafpoi))
	for i, raf := range rafpoi {
		ok := false
		for selrafsi, rafsiporsi := range jz.Rafsi {
      if (len(raf) == 4 || len(raf) == 5) && len(selrafsi) == 5 && selrafsi[:4] == string(raf[:4]) {
				tanru[i] = selrafsi
        ok = true
      } else {
        for _, rafsi := range rafsiporsi {
          if rafsi == string(raf) {
            tanru[i] = selrafsi
            ok = true
            break
          }
        }
      }
			if ok {
				break
			}
		}
		if !ok {
			if len(raf) > 5 || (len(raf) == 5 && (bytes.Contains([]byte("aeiou"), []byte{raf[0]}) || bytes.Contains([]byte("aeiou"), []byte{raf[3]}))) || (len(raf) == 4 && bytes.Contains([]byte("aeiou"), []byte{raf[0]}) && !bytes.Contains([]byte("aeiou"), []byte{raf[1]}) && !bytes.Contains([]byte("aeiou"), []byte{raf[2]}) && bytes.Contains([]byte("aeiou"), []byte{raf[3]})) {
				tanru[i] = string(raf)
			} else {
				respond("unknown rafsi -" + string(raf) + "-")
				return
			}
		}
	}
	respond(strings.Join(tanru, " "))
}
