package actions

import (
	"fmt"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

func Edit(scope flags.Scope, value *string) {
	result := config.GetContentByEditor("Stuff to edit for now.")
	fmt.Println(result)
}
