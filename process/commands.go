package process

import (
	"fmt"

	"github.com/ducktordanny/td/actions"
	"github.com/ducktordanny/td/flags"
)

func CommandsToAction(flagValues *flags.FlagValues) {
	fmt.Println(*flagValues)

	scope := flags.Scope(flags.LocalScope)
	if flagValues.Global == true {
		scope = flags.Scope(flags.GlobalScope)
	}

	if flagValues.Add != "" {
		actions.Add(scope, &flagValues.Add)
	} else if flagValues.Remove != "" {
		actions.Remove(scope, &flagValues.Remove)
	} else if flagValues.Resolve != "" {
		actions.Resolve(scope, &flagValues.Resolve)
	} else if flagValues.Unresolve != "" {
		actions.Unresolve(scope, &flagValues.Unresolve)
	} else if flagValues.Toggle != "" {
		actions.Toggle(scope, &flagValues.Toggle)
	} else if flagValues.Edit != false {
		actions.Edit(scope, &flagValues.Edit)
	} else if flagValues.List != false {
		actions.List(scope, &flagValues.List)
	} else if flagValues.ListResolved != false {
		actions.ListResolved(scope, &flagValues.ListResolved)
	} else if flagValues.ListUnresolved != false {
		actions.ListUnresolved(scope, &flagValues.ListUnresolved)
	}
}
