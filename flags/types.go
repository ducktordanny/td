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
	Edit           string
	List           bool
	ListResolved   bool
	ListUnresolved bool
}
