package config

type TodoModel struct {
	content  string
	resolved bool
}

type LocalsMap []map[string][]TodoModel

type TdConfig struct {
	globals []TodoModel
	locals  LocalsMap
}
