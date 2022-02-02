package gcg

type FileData struct {
	Package      string
	BuildVersion string
	Types        []TypeData
	Enums        []EnumData
	Mutations    []OperationData
	Queries      []OperationData
	Scalars      []ScalarData
	Functions    []FunctionData
	Unions       []UnionData
}

type Argument struct {
	Name     string
	JSONName string
	Type     string
	IsArray  bool
}

type OperationData struct {
	Comment        string
	Name           string
	FunctionName   string
	Arguments      []Argument
	ReturnType     string
	ReturnFieldTag string
	IsArray        bool
	DataTag        string
	ErrorsTag      string
}

type TypeData struct {
	Comment string
	Kind    string
	Name    string
	Fields  []FieldData
}

type FieldData struct {
	Comment  string
	Name     string
	Tag      string
	Type     string
	IsArray  bool
	Optional bool
}

type EnumData struct {
	Comment string
	Name    string
	Values  []string
}

type ScalarData struct {
	Comment string
	Name    string
}

type FunctionData struct {
	Comment    string
	Name       string
	Fields     []FieldData
	IsArray    bool
	ReturnType string
	ResultTag  string
}

type UnionData struct {
	Comment string
	Name    string
	Types   []string
}
