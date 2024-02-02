package main

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func main() {
	flagValues := flags.Init()
	fmt.Println(flagValues)
}
