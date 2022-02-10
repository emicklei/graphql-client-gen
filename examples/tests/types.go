package tests

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg version: (dev)
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

// Identified is a INTERFACE.
type Identified struct {
	ID interface{} `graphql:"id" json:"id,omitempty"`
}

// Result is a OBJECT.
type Result struct {
	When *CustomDate `graphql:"when" json:"when,omitempty"`
}

// ResultInput is a INPUT_OBJECT.
type ResultInput struct {
	When   CustomDate `graphql:"when" json:"when,omitempty"`
	Unused *string    `graphql:"unused" json:"unused,omitempty"`
}
