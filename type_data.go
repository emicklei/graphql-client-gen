package gcg

import "time"

type FileData struct {
	Package    string
	Created    time.Time
	Types      []TypeData
	Enums      []EnumData
	Operations []OperationData
}

type Argument struct {
	Name    string
	Type    string
	IsArray bool
}

type OperationData struct {
	Name         string
	FunctionName string
	Arguments    []Argument
	ReturnType   string
	IsArray      bool
	Tag          string
}

type TypeData struct {
	Kind          string
	EmbeddedTypes []string
	Name          string
	Fields        []FieldData
}

type FieldData struct {
	Name     string
	Tag      string
	Type     string
	IsArray  bool
	Optional bool
}

type EnumData struct {
	Name   string
	Values []string
}
