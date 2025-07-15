package tests

// copied from build_source.go by github.com/emicklei/graphql-client-gen/cmd/gcg version: v1.0.1+dirty
// DO NOT EDIT

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

func buildRequest(mutationOrQuery string, operationName string, querySample any, typedVars map[string]valueAndType) GraphQLRequest {
	queryBody := new(bytes.Buffer)
	writeQuery(querySample, queryBody, 0, false, typedVars)
	body := queryBody.String()
	signature := new(bytes.Buffer)
	keys := []string{}
	for k := range typedVars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := typedVars[k]
		if signature.Len() > 0 {
			io.WriteString(signature, ",")
		}
		fmt.Fprintf(signature, "$%s:%v", k, v.graphType)
	}
	vars := map[string]any{}
	for k, v := range typedVars {
		vars[k] = v.value
	}
	return GraphQLRequest{
		Query:         fmt.Sprintf("%s %s(%s) {\n%s}", mutationOrQuery, operationName, signature.String(), body),
		OperationName: operationName,
		Variables:     vars,
	}
}

func writeQuery(q any, w io.Writer, indent int, inline bool, vars map[string]valueAndType) {
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
	isDataRoot := indent == 0
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if fv.IsZero() && !isDataRoot {
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
			if k.Kind() == reflect.Slice && fv.Len() > 0 {
				writeQuery(fv.Index(0).Interface(), w, indent, inline, vars)
			} else if k.Kind() == reflect.Struct {
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

func writeStruct(structValue any, w io.Writer, indent int, inline bool, vars map[string]valueAndType) {
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
	isDataRoot := indent == 0
	// do not write empty nested structure if no fields are requested
	// unless data root which describes the query | mutation
	if isZeroGraphQLStruct(reflect.ValueOf(structValue)) && !isDataRoot {
		io.WriteString(w, "\n")
		return
	}
	io.WriteString(w, " {\n")
	io.WriteString(w, strings.Repeat("\t", indent+1))
	writeQuery(structValue, w, indent+1, inline, vars)
	io.WriteString(w, "}\n")
}

func isZeroGraphQLStruct(v reflect.Value) bool {
	rt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := rt.Field(i)
		if len(f.Tag) == 0 { // no tag at all
			continue
		}
		_, ok := f.Tag.Lookup("graphql-function-arg") // ignore arguments
		if ok {
			continue
		}
		if !v.Field(i).IsZero() {
			return false
		}
	}
	return true
}

type valueAndType struct {
	name      string
	value     any
	graphType string
}

func collectionFunctionArgs(structValue any) (list []valueAndType) {
	rt := reflect.TypeOf(structValue)
	rv := reflect.ValueOf(structValue)
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if fv.IsZero() {
			continue
		}
		ft := rt.Field(i)
		tag, ok := ft.Tag.Lookup("graphql-function-arg")
		if ok {
			gt, _ := ft.Tag.Lookup("graphql-function-type")
			list = append(list, valueAndType{name: tag, value: fv.Interface(), graphType: gt})
		}
	}
	return
}

// GraphQLRequest is used to model both a query or a mutation request
type GraphQLRequest struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operationName"`
	Variables     map[string]any `json:"variables"`
}

// NewGraphQLRequest returns a new Request (for query or mutation) with optional or empty variables.
func NewGraphQLRequest(query, operation string, vars ...map[string]any) GraphQLRequest {
	initVars := map[string]any{}
	if len(vars) > 0 {
		initVars = vars[0] // merge all?
	}
	return GraphQLRequest{Query: query, OperationName: operation, Variables: initVars}
}

// Error is a response field element to capture server reported problems
type Error struct {
	Message   string `json:"message,omitempty"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
}
