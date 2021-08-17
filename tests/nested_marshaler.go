package tests

import (
	"github.com/CosmWasm/tinyjson"
	"github.com/CosmWasm/tinyjson/jlexer"
	"github.com/CosmWasm/tinyjson/jwriter"
)

//tinyjson:json
type NestedMarshaler struct {
	Value  tinyjson.MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalTinyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalTinyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}
