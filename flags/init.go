package flags

import (
	"flag"
	"fmt"
)

func getAliasUsage(name string) string {
	return fmt.Sprintf("Alias for '%s'.", name)
}

func initStringVarFlag(props *FlagProperties[string]) {
	aliasUsage := getAliasUsage(props.name)
	flag.StringVar(&props.value, props.alias, props.defaultValue, aliasUsage)
	flag.StringVar(&props.value, props.name, props.defaultValue, props.usage)
}

func initBoolVarFlag(props *FlagProperties[bool]) {
	aliasUsage := getAliasUsage(props.name)
	flag.BoolVar(&props.value, props.alias, props.defaultValue, aliasUsage)
	flag.BoolVar(&props.value, props.name, props.defaultValue, props.usage)
}

func Init() FlagValues {
	initStringVarFlag(&Add)
	initStringVarFlag(&Remove)
	initStringVarFlag(&Resolve)
	initStringVarFlag(&Unresolve)
	initStringVarFlag(&Toggle)

	initBoolVarFlag(&Edit)
	initBoolVarFlag(&List)
	initBoolVarFlag(&ListResolved)
	initBoolVarFlag(&ListUnresolved)

	flag.Parse()

	return FlagValues{
		Add:            Add.value,
		Remove:         Remove.value,
		Resolve:        Resolve.value,
		Unresolve:      Unresolve.value,
		Toggle:         Toggle.value,
		Edit:           Edit.value,
		List:           List.value,
		ListResolved:   ListResolved.value,
		ListUnresolved: ListUnresolved.value,
	}
}
