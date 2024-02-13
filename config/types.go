package config

type TodoModel struct {
	Id       string
	Content  string
	Resolved bool
}

type LocalsMap map[string][]TodoModel

type TdConfig struct {
	Globals []TodoModel
	Locals  LocalsMap
}
