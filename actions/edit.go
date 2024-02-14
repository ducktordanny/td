package actions

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func Edit(scope flags.Scope, value *bool) {
	fmt.Println("edit", *value, scope)
	// TODO: implement
}
