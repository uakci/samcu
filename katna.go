package main

import (
  jz "github.com/uakci/jvozba/v3"
  "bytes"
  "strings"
)

func katna(respond func(string), _ string, args []string) {
	if len(args) != 1 {
		respond("one argument expected")
		return
	}
	arg := h.Replace(args[0])
  decomp := cutter(arg)
  for i, a := range decomp {
    if a == "" {
      respond("rafsi -" + string(jz.Katna([]byte(arg))[i]) + "- unknown")
      return
    }
  }
  respond(strings.Join(decomp, " "))
}

func cutter(lujvo string) []string {
	rafpoi := jz.Katna([]byte(lujvo))
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
				tanru[i] = ""
			}
		}
	}
  return tanru
}
