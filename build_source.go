package gcg

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func BuildQuery(q interface{}) string {
	b := new(bytes.Buffer)
	writeQuery(q, b, 0, false)
	return b.String()
}

func writeQuery(q interface{}, w io.Writer, indent int, inline bool) {
	rt := reflect.TypeOf(q)
	rv := reflect.ValueOf(q)
	if rt.Kind() == reflect.Ptr {
		writeQuery(rv.Elem().Interface(), w, indent, inline)
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if !fv.IsZero() {
			sf := rt.Field(i)
			tag, ok := sf.Tag.Lookup("graphql")
			inlineField := sf.Anonymous && !ok
			if inlineField {
				// is struct or pointer to struct
				k := sf.Type
				if k.Kind() == reflect.Ptr {
					k = k.Elem()
					fv = fv.Elem()
				}
				if k.Kind() == reflect.Struct {
					writeQuery(fv.Interface(), w, indent, inlineField)
				}
			} else {
				// no inline
				if ok {
					// handle operation override for query
					if overrides, ok := q.(OverridesOperationName); ok {
						op := overrides.OperationName()
						if op == "" {
							op = "unsetOperation"
						}
						tag = strings.Replace(tag, "OperationName", op, -1)
					}
					fmt.Fprintf(w, "\t%s", tag)
					// is struct or pointer to struct
					k := sf.Type
					if k.Kind() == reflect.Ptr {
						k = k.Elem()
						fv = fv.Elem()
					}
					if k.Kind() == reflect.Struct {
						writeStruct(fv.Interface(), w, indent, inlineField)
					} else if k.Kind() == reflect.Slice {
						// take first if avail
						if fv.Len() > 0 {
							// always struct? TODO
							writeStruct(fv.Index(0).Interface(), w, indent, inline)
						}
					} else {
						io.WriteString(w, "\n")
					}
					io.WriteString(w, strings.Repeat("\t", indent))
				}
			}
		}
	}
}

func writeStruct(structValue interface{}, w io.Writer, indent int, inline bool) {
	if list := collectionFunctionArgs(structValue); len(list) > 0 {
		io.WriteString(w, "(")
		for i, each := range list {
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprintf(w, "%s:%v", each.k, each.v)
		}
		io.WriteString(w, ")")
	}
	io.WriteString(w, " {\n")
	io.WriteString(w, strings.Repeat("\t", indent+1))
	writeQuery(structValue, w, indent+1, inline)
	io.WriteString(w, "}\n")
}

type kv struct {
	k string
	v interface{}
}

func collectionFunctionArgs(structValue interface{}) (list []kv) {
	rt := reflect.TypeOf(structValue)
	rv := reflect.ValueOf(structValue)
	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)
		if !fv.IsZero() {
			tag, ok := ft.Tag.Lookup("graphql-function-arg")
			if ok {
				list = append(list, kv{k: tag, v: fv.Interface()})
			}
		}
	}
	return
}

// GraphQLRequest is used to model both a query or a mutation request
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// OverridesOperationName is used to replace the "OperationName" in the tag when building a GraphQL query
type OverridesOperationName interface {
	OperationName() string
}

type Errors struct {
	Message   string `json:"message,omitempty"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}
