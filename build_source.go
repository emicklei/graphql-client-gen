package gcg

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func buildRequest(mq string, op string, q interface{}, typedVars map[string]valueAndType) GraphQLRequest {
	b := new(bytes.Buffer)
	writeQuery(q, b, 0, false, typedVars)
	body := b.String()
	s := new(bytes.Buffer)
	for k, v := range typedVars {
		if s.Len() > 0 {
			io.WriteString(s, ",")
		}
		fmt.Fprintf(s, "$%s:%v", k, v.graphType)
	}
	vars := map[string]interface{}{}
	for k, v := range typedVars {
		vars[k] = v.value
	}
	return GraphQLRequest{
		Query:         fmt.Sprintf("%s %s(%s)\n{%s\n}", mq, op, s.String(), body),
		OperationName: op,
		Variables:     vars,
	}
}

func writeQuery(q interface{}, w io.Writer, indent int, inline bool, vars map[string]valueAndType) {
	rt := reflect.TypeOf(q)
	rv := reflect.ValueOf(q)
	if rt.Kind() == reflect.Ptr {
		writeQuery(rv.Elem().Interface(), w, indent, inline, vars)
		return
	}
	if rt.Kind() == reflect.Slice {
		// take first if avail
		if rv.Len() > 0 {
			writeQuery(rv.Index(0).Interface(), w, indent, inline, vars)
		}
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if fv.IsZero() {
			continue
		}
		sf := rt.Field(i)
		tag, ok := sf.Tag.Lookup("graphql")
		inlineField := tag == "inline" || (sf.Anonymous && !ok)
		if inlineField {
			// is struct or pointer to struct
			k := sf.Type
			if k.Kind() == reflect.Ptr {
				k = k.Elem()
				fv = fv.Elem()
			}
			if k.Kind() == reflect.Struct {
				writeQuery(fv.Interface(), w, indent, inlineField, vars)
			}
		} else {
			// no inline
			if ok {
				fmt.Fprintf(w, "\t%s", tag)
				writeField(sf, fv, w, indent, false, vars)
			}
		}
	}
}

func writeField(sf reflect.StructField, fv reflect.Value, w io.Writer, indent int, inline bool, vars map[string]valueAndType) {
	// is struct or pointer to struct
	k := sf.Type
	if k.Kind() == reflect.Ptr {
		k = k.Elem()
		fv = fv.Elem()
	}
	if k.Kind() == reflect.Struct {
		writeStruct(fv.Interface(), w, indent, false, vars)
	} else if k.Kind() == reflect.Slice {
		// take first if avail
		if fv.Len() > 0 {
			// always struct? TODO
			writeStruct(fv.Index(0).Interface(), w, indent, inline, vars)
		}
	} else {
		io.WriteString(w, "\n")
	}
	io.WriteString(w, strings.Repeat("\t", indent))
}

func writeStruct(structValue interface{}, w io.Writer, indent int, inline bool, vars map[string]valueAndType) {
	if list := collectionFunctionArgs(structValue); len(list) > 0 {
		io.WriteString(w, "(")
		for i, each := range list {
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			varName := fmt.Sprintf("%s%d", each.name, len(vars))
			fmt.Fprintf(w, "%s:$%s", each.name, varName)
			vars[varName] = each
		}
		io.WriteString(w, ")")
	}
	// do not write empty nested structure if no fields are requested
	if reflect.ValueOf(structValue).IsZero() {
		io.WriteString(w, "\n")
		return
	}
	io.WriteString(w, " {\n")
	io.WriteString(w, strings.Repeat("\t", indent+1))
	writeQuery(structValue, w, indent+1, inline, vars)
	io.WriteString(w, "}\n")
}

type valueAndType struct {
	name      string
	value     interface{}
	graphType string
}

func collectionFunctionArgs(structValue interface{}) (list []valueAndType) {
	rt := reflect.TypeOf(structValue)
	rv := reflect.ValueOf(structValue)
	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)
		if !fv.IsZero() {
			tag, ok := ft.Tag.Lookup("graphql-function-arg")
			if ok {
				gt, _ := ft.Tag.Lookup("graphql-function-type")
				list = append(list, valueAndType{name: tag, value: fv.Interface(), graphType: gt})
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

// Errors is a response field to capture server reported problems
type Errors struct {
	Message   string `json:"message,omitempty"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}
