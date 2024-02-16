package config

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
