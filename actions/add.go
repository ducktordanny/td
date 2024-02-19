package actions

import (
	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

func Add(scope flags.Scope, value *string) {
	todo := config.TodoModel{
		Id:       config.GenerateUniqueSha(),
		Content:  *value,
		Resolved: false,
	}
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	*list = append(*list, todo)
	config.WriteConfig(conf)
}
