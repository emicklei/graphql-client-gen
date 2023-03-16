package gcg

type FileData struct {
	Package         string
	BuildVersion    string
	Types           []TypeData
	Enums           []EnumData
	Mutations       []OperationData
	Queries         []OperationData
	Scalars         []ScalarData
	Functions       []FunctionData
	Unions          []UnionData
	Inputs          []InputData
	IncludeScalarID bool
	SchemaVersion   string
}

type Argument struct {
	Name        string
	JSONName    string
	Type        string
	GraphType   string
	IsArray     bool
	ArraySuffix string
}

type OperationData struct {
	Comment         string
	Definition      string
	Name            string
	FunctionName    string
	Arguments       []Argument
	ReturnType      string
	ReturnFieldName string
	EmbedFieldTag   string
	ReturnFieldTag  string
	IsArray         bool
	ArraySuffix     string
	DataTag         string
	ErrorsTag       string
}

type TypeData struct {
	Comment     string
	Kind        string
	Name        string
	Fields      []FieldData
	TypenameTag string
}

type FieldData struct {
	StructName  string
	Comment     string
	Name        string
	JSONName    string
	Tag         string
	Type        string
	GraphType   string
	IsArray     bool
	ArraySuffix string
	Optional    bool
	Deprecated  bool
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
	Comment     string
	Name        string
	Fields      []FieldData
	IsArray     bool
	ArraySuffix string
	ReturnType  string
	ResultTag   string
}

type UnionData struct {
	Comment string
	Name    string
	Types   []string
}

type InputData struct {
	Comment string
	Name    string
	Fields  []FieldData
}
