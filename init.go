package samcu

import (
	"fmt"
	"github.com/uakci/samcu/jvs"
	"os"
)

func Init() {
	cookieLine, ok := os.LookupEnv("JVS_COOKIE")
  if !ok {
    panic(fmt.Errorf("Need %t in env var %t", "token", "JVS_COOKIE"))
  }

  okChan := make(chan struct{})
	go func() {
    jvs.FetchAll(cookieLine, Index, okChan)
  }()
  <-okChan
}

const defaultLang = "en"

var AllCommands = map[string]Command{
  "ciksi": Ciksi,
  "gerna": Gerna,
  "katna": Katna,
  "lujvo": Lujvo,
  "rafsi": Rafsi,
  "sisku": Sisku,
}

type Command struct{
  Name string
  Func func(map[string]any) (string, error)
  Desc string
  Args []CommandOption
  Opts []CommandOption
}

type CommandOption struct{
  Name string
  Desc string
  Values [][2]string
  Type OptionType
}

type OptionType int8
const (
  StringType OptionType = iota
  BoolType
)

var BanguOpt = CommandOption{"bangu", "vlaste bangu", nil, StringType}

func WithDefault[T any](args map[string]any, key string, def T) T {
  if val, ok := args[key]; ok {
    return val.(T)
  } else {
    return def
  }
}

func GetBangu(args map[string]any) string {
  if bangu, ok := args["bangu"]; ok {
    return bangu.(string)
  } else {
    return defaultLang
  }
}
