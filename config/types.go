package config

// TODO: It would be good to have the date of creation
type TodoModel struct {
	Id       string
	Content  string
	Resolved bool
}

type LocalModel struct {
	Path  string
	Items []TodoModel
}

type TdConfig struct {
	Globals []TodoModel
	Locals  []LocalModel
}
