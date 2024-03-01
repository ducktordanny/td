package flags

var Add = FlagProperties[bool]{
	usage: "Creates a new TODO.",
	name:  "add",
	alias: "a",
}

var Remove = FlagProperties[string]{
	usage: "Removes a TODO based on the provided ID.",
	name:  "remove",
	alias: "rm",
}

var Resolve = FlagProperties[string]{
	usage: "Resolves a TODO based on the provided ID.",
	name:  "resolve",
	alias: "rs",
}

var Unresolve = FlagProperties[string]{
	usage: "Unresolves a TODO based on the provided ID.",
	name:  "unresolve",
	alias: "urs",
}

var Toggle = FlagProperties[string]{
	usage: "Toggles the resolved status of a TODO based on the provided ID.",
	name:  "toggle",
	alias: "t",
}

var Edit = FlagProperties[string]{
	usage: "Edits a TODO based on the provided ID.",
	name:  "edit",
	alias: "e",
}

var List = FlagProperties[bool]{
	usage: "Returns the list of all TODOs in the scope.",
	name:  "list",
	alias: "ls",
}

var ListResolved = FlagProperties[bool]{
	usage: "Returns the list of resolved TODOs in the scope.",
	name:  "list-resolved",
	alias: "ls-rs",
}

var ListUnresolved = FlagProperties[bool]{
	usage: "Returns the list of unresolved TODOs in the scope.",
	name:  "list-unresolved",
	alias: "ls-urs",
}

var Global = FlagProperties[bool]{
	usage: "Switches to global scope, and handles global TODOs.",
	name:  "global",
	alias: "g",
}
