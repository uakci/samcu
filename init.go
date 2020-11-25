package samcu

import (
	"encoding/json"
	"os"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	myRole      = "772142260128710719"
	helpChannel = "772167961771245578"
)

func LoadJVS() error {
	f, e := os.Open("jvs.json")
	if e != nil {
		return e
	}
	e = json.NewDecoder(f).Decode(&dict)
	return e
}
