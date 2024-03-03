package flags

import (
	"flag"
	"fmt"
)

func getAliasUsage(name string) string {
	return fmt.Sprintf("Alias for '%s'.", name)
}

func initStringVarFlag(props *FlagProperties[string]) {
	if props.alias != "" {
		aliasUsage := getAliasUsage(props.name)
		flag.StringVar(&props.value, props.alias, props.defaultValue, aliasUsage)
	}
	flag.StringVar(&props.value, props.name, props.defaultValue, props.usage)
}

func initBoolVarFlag(props *FlagProperties[bool]) {
	aliasUsage := getAliasUsage(props.name)
	flag.BoolVar(&props.value, props.alias, props.defaultValue, aliasUsage)
	flag.BoolVar(&props.value, props.name, props.defaultValue, props.usage)
}

func Init() FlagValues {
	initBoolVarFlag(&Add)
	initStringVarFlag(&Remove)
	initStringVarFlag(&Resolve)
	initStringVarFlag(&Unresolve)
	initStringVarFlag(&Toggle)
	initStringVarFlag(&Edit)
	initStringVarFlag(&Content)
	initBoolVarFlag(&List)
	initBoolVarFlag(&ListResolved)
	initBoolVarFlag(&ListUnresolved)
	initBoolVarFlag(&Global)
	flag.Parse()

	return FlagValues{
		Add:            Add.value,
		Remove:         Remove.value,
		Resolve:        Resolve.value,
		Unresolve:      Unresolve.value,
		Toggle:         Toggle.value,
		Edit:           Edit.value,
		Content:        Content.value,
		List:           List.value,
		ListResolved:   ListResolved.value,
		ListUnresolved: ListUnresolved.value,
		Global:         Global.value,
	}
}
