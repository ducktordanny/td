package config

import "time"

type TodoModel struct {
	Id         string
	Content    string
	Resolved   bool
	CreatedAt  time.Time
	ResolvedAt *time.Time `json:"ResolvedAt,omitempty"`
}

type LocalModel struct {
	Path  string
	Items []TodoModel
}

type TdConfig struct {
	Globals []TodoModel
	Locals  []LocalModel
}
