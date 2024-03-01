package actions

import (
	"log"
	"time"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
)

func Edit(scope flags.Scope, value *string) {
	if len(*value) != 7 {
		log.Fatalln("Invalid Sha ID.")
	}
	conf := config.ReadConfig()
	list := config.GetListBasedOnScope(scope, &conf)
	indexToEdit := config.GetIndexOfListBySha(*list, *value)
	if indexToEdit == -1 {
		log.Fatalln("Invalid Sha ID.")
	}
	result := config.GetContentByEditor((*list)[indexToEdit].Content)
	(*list)[indexToEdit].Content = result
	(*list)[indexToEdit].CreatedAt = time.Now()
	config.WriteConfig(conf)
}
