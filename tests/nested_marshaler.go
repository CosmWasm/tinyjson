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

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}
