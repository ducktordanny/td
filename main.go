package main

import (
	"github.com/ducktordanny/td/flags"
	"github.com/ducktordanny/td/process"
)

func main() {
	flagValues := flags.Init()
	process.CommandsToAction(&flagValues)
}
