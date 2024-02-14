package actions

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func Add(scope flags.Scope, value *string) {
	fmt.Println("add", *value, scope)
}
