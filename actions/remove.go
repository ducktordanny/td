package actions

import (
	"log"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
	"golang.org/x/exp/slices"
)

// TODO: option for remove last and first
func Remove(scope flags.Scope, value *string) {
	if len(*value) != 7 {
		log.Fatalln("Invalid Sha ID.")
	}
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	indexToRemove := config.GetIndexOfListBySha(*list, *value)
	if indexToRemove == -1 {
		log.Fatalln("Invalid Sha ID.")
	}
	*list = slices.Delete[[]config.TodoModel, config.TodoModel](*list, indexToRemove, indexToRemove+1)
	config.WriteConfig(conf)
}
