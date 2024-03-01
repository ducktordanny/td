package actions

import (
	"fmt"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

type resolveFilter string

const (
	resolved   resolveFilter = "resolved"
	unresolved resolveFilter = "unresolved"
	none       resolveFilter = ""
)

func formatToString(list []config.TodoModel, filter resolveFilter) {
	const (
		Yellow = "\033[33m"
		Reset  = "\033[0m"
	)
	for index, item := range list {
		if filter == resolved && !item.Resolved {
			continue
		}
		if filter == unresolved && item.Resolved {
			continue
		}
		output := "%stodo %s%s\nDate: %s\nResolved: %t\n\n%s\n"
		if index < len(list)-1 {
			output = output + "\n"
		}
		fmt.Printf(output, Yellow, item.Id, Reset, item.CreatedAt, item.Resolved, item.Content)
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
