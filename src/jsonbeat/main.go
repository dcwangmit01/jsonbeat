package main

import (
	"os"
	"github.com/elastic/beats/libbeat/beat"
	"jsonbeat/beater"
)

var Name = "jsonbeat"

func main() {
	if err := beat.Run(Name, "0.0.1", beater.New()); err != nil {
		os.Exit(1)
	}
}
