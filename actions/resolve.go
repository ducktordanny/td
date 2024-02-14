package actions

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func Resolve(scope flags.Scope, value *string) {
	fmt.Println("resolve", *value, scope)
	// TODO: implement
}

func Unresolve(scope flags.Scope, value *string) {
	fmt.Println("unresolve", *value, scope)
	// TODO: implement
}

func Toggle(scope flags.Scope, value *string) {
	fmt.Println("toggle", *value, scope)
	// TODO: implement
}
