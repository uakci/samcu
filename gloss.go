package main

import (
  jz "github.com/uakci/jvozba/v3"
  "strings"
)

func gloss(respond func(string), _ string, args []string) {
  result := strings.Builder{}
  for i, a := range args {
    if i > 0 {
      result.WriteRune('\u2003')
    }
    a = h.Replace(a)
    var subject []string
    if jz.IsGismu([]byte(a)) || jz.IsCmavo([]byte(a)) || len(jz.Katna([]byte(a))) == 1 {
      subject = []string{a}
    } else {
      subject = cutter(a)
    }
    for j, s := range subject {
      if j > 0 {
        result.WriteRune('—')
      }
      def, ok := dict[lang][s]
      if !ok || len(def.Glosses) == 0 {
        result.WriteString("*" + s + "*")
      } else {
        result.WriteString(def.Glosses[0])
      }
    }
  }
  respond(result.String())
}
