package actions

import (
	"fmt"

	"github.com/ducktordanny/td/flags"
)

func List(scope flags.Scope, value *bool) {
	fmt.Println("list", *value, scope)
	// TODO: implement
}

func ListResolved(scope flags.Scope, value *bool) {
	fmt.Println("list resolved", *value, scope)
	// TODO: implement
}

func ListUnresolved(scope flags.Scope, value *bool) {
	fmt.Println("list unresolved", *value, scope)
	// TODO: implement
}
