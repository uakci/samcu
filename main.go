package samcu

import (
	"fmt"
	"os"

	"github.com/uakci/samcu/jvs"
)

const defaultBangu = "en"

func Init() {
	cookieLine, ok := os.LookupEnv("JVS_COOKIE")
	if !ok {
		panic(fmt.Errorf("Need %s in env var %s", "token", "JVS_COOKIE"))
	}

	okChan := make(chan struct{})
	go func() {
		jvs.FetchAll(cookieLine, Index, okChan)
	}()
	<-okChan
}

type Handler func([]string) (string, error)

var Commands = map[string]Handler{
	"ciksi":   WithDictionary(ciksi),
	"gerna":   NoEmpty(gerna),
	"katna":   Repeatedly(katna),
	"lujvo":   NoEmpty(lujvo),
	"marafsi": Repeatedly(marafsi),
	"rafsima": Repeatedly(rafsima),
	"sisku":   WithDictionary(sisku),
}

func Respond(args []string) (ok bool, msg string, err error) {
	ok, msg, err = false, "", nil
	if len(args) == 0 {
		return
	}

	var handler Handler
	handler, ok = Commands[args[0]]
	if !ok {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("u'u spaji nabmi se cu'u la'e di'e fa'o: %v", r)
		}
	}()

	msg, err = handler(args[1:])
	return
}
