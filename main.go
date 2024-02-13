package main

import (
	"fmt"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
	"github.com/ducktordanny/td/process"
)

func main() {
	flagValues := flags.Init()
	process.CommandsToAction(&flagValues)
	test := config.ReadConfig()
	fmt.Println(test)
}
