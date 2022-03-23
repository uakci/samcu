package common

import "strings"

var h = strings.NewReplacer("’", "'", "h", "'", ".", "")

func ReplaceH(s string) string {
  return h.Replace(s)
}
