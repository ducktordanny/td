package flags

type FlagProperties[T any] struct {
	value        T
	defaultValue T
	usage        string
	name         string
	alias        string
}

type FlagValues struct {
	Add            string
	Remove         string
	Resolve        string
	Unresolve      string
	Toggle         string
	Edit           bool
	List           bool
	ListResolved   bool
	ListUnresolved bool
}
