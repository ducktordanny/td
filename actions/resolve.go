package actions

import (
	"log"
	"time"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

type resolveCallback func(list *[]config.TodoModel, indexToResolve int)

func resolve(scope flags.Scope, value *string, cb resolveCallback) {
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	indexToResolve := config.GetIndexOfListBySha(*list, *value)
	if indexToResolve == -1 {
		log.Fatalln("Invalid Sha ID.")
	}
	cb(list, indexToResolve)
	config.WriteConfig(conf)
}

func Resolve(scope flags.Scope, value *string) {
	resolve(scope, value, func(list *[]config.TodoModel, indexToResolve int) {
		now := time.Now()
		(*list)[indexToResolve].ResolvedAt = &now
		(*list)[indexToResolve].Resolved = true
	})
}

func Unresolve(scope flags.Scope, value *string) {
	resolve(scope, value, func(list *[]config.TodoModel, indexToResolve int) {
		(*list)[indexToResolve].ResolvedAt = nil
		(*list)[indexToResolve].Resolved = false
	})
}

func Toggle(scope flags.Scope, value *string) {
	resolve(scope, value, func(list *[]config.TodoModel, indexToResolve int) {
		isResolved := (*list)[indexToResolve].Resolved
		if isResolved {
			(*list)[indexToResolve].ResolvedAt = nil
		} else {
			now := time.Now()
			(*list)[indexToResolve].ResolvedAt = &now
		}
		(*list)[indexToResolve].Resolved = !isResolved
	})
}
