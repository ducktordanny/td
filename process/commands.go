package process

import (
	"fmt"

	"github.com/ducktordanny/td/actions"
	"github.com/ducktordanny/td/flags"
)

func CommandsToAction(flagValues *flags.FlagValues) {
	fmt.Println(*flagValues)

	if flagValues.Add != "" {
		actions.Add(&flagValues.Add)
	} else if flagValues.Remove != "" {
		actions.Remove(&flagValues.Remove)
	} else if flagValues.Resolve != "" {
		actions.Resolve(&flagValues.Resolve)
	} else if flagValues.Unresolve != "" {
		actions.Unresolve(&flagValues.Unresolve)
	} else if flagValues.Toggle != "" {
		actions.Toggle(&flagValues.Toggle)
	} else if flagValues.Edit != false {
		actions.Edit(&flagValues.Edit)
	} else if flagValues.List != false {
		actions.List(&flagValues.List)
	} else if flagValues.ListResolved != false {
		actions.ListResolved(&flagValues.ListResolved)
	} else if flagValues.ListUnresolved != false {
		actions.ListUnresolved(&flagValues.ListUnresolved)
	}
}
