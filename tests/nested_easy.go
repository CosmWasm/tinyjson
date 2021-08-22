package tests

import (
	"github.com/CosmWasm/tinyjson"
	"github.com/CosmWasm/tinyjson/jwriter"
)

//tinyjson:json
type NestedInterfaces struct {
	Value interface{}
	Slice []interface{}
	Map   map[string]interface{}
}

type NestedEasyMarshaler struct {
	EasilyMarshaled bool
}

var _ tinyjson.Marshaler = &NestedEasyMarshaler{}

func (i *NestedEasyMarshaler) MarshalTinyJSON(w *jwriter.Writer) {
	// We use this method only to indicate that tinyjson.Marshaler
	// interface was really used while encoding.
	i.EasilyMarshaled = true
}
