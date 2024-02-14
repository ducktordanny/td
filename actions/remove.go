package actions

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func Remove(scope flags.Scope, value *string) {
	fmt.Println("remove", *value, scope)
	// TODO: implement
}
