package actions

import (
	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

func getListBasedOnScope(scope flags.Scope, conf *config.TdConfig) *[]config.TodoModel {
	if scope == flags.LocalScope {
		path := config.GetLocalPath()
		localModel := config.GetLocalModelByPath(conf, path)
		return &localModel.Items
	}
	return &(conf.Globals)
}

func Add(scope flags.Scope, value *string) {
	todo := config.TodoModel{
		Id:       config.GenerateUniqueSha(),
		Content:  *value,
		Resolved: false,
	}
	conf := config.ReadFullConfig()
	list := getListBasedOnScope(scope, &conf)
	*list = append(*list, todo)
	config.WriteFullConfig(conf)
}
