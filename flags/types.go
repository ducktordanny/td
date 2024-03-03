package flags

type Scope string

const (
	LocalScope  Scope = "local"
	GlobalScope Scope = "global"
)

type FlagProperties[T any] struct {
	value        T
	defaultValue T
	usage        string
	name         string
	alias        string
}

type FlagValues struct {
	Add            bool
	Remove         string
	Resolve        string
	Unresolve      string
	Toggle         string
	Edit           string
	List           bool
	ListResolved   bool
	ListUnresolved bool
	Global         bool
	Content        string
}
