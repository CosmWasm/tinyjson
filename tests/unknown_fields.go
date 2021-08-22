package tests

import "github.com/CosmWasm/tinyjson"

//tinyjson:json
type StructWithUnknownsProxy struct {
	tinyjson.UnknownFieldsProxy

	Field1 string
}

//tinyjson:json
type StructWithUnknownsProxyWithOmitempty struct {
	tinyjson.UnknownFieldsProxy

	Field1 string `json:",omitempty"`
}
