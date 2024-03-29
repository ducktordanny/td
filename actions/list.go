package actions

import (
	"fmt"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
	"github.com/fatih/color"
)

type resolveFilter string

const (
	resolved   resolveFilter = "resolved"
	unresolved resolveFilter = "unresolved"
	none       resolveFilter = ""
)

func formatToString(list []config.TodoModel, filter resolveFilter) {
	for index, item := range list {
		if filter == resolved && !item.Resolved {
			continue
		}
		if filter == unresolved && item.Resolved {
			continue
		}

		idOutput := "todo %s\n"
		output := "Date: %s\nResolved: %t\n\n%s\n"
		// Handle extra linebreak if it's not the last item
		if index < len(list)-1 {
			output = output + "\n"
		}

		color.Set(color.FgYellow)
		fmt.Printf(idOutput, item.Id)
		color.Unset()
		fmt.Printf(output, item.CreatedAt, item.Resolved, item.Content)
	}
}

func List(scope flags.Scope, value *bool) {
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	formatToString(*list, none)
}

func ListResolved(scope flags.Scope, value *bool) {
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	formatToString(*list, resolved)
}

func ListUnresolved(scope flags.Scope, value *bool) {
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	formatToString(*list, unresolved)
}
