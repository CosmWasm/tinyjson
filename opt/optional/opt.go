// +build none

package optional

import (
	"fmt"

	"github.com/CosmWasm/tinyjson/jlexer"
	"github.com/CosmWasm/tinyjson/jwriter"
)

// template type Optional(A)
type A int

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Optional struct {
	V       A
	Defined bool
}

// Creates an optional type with a given value.
func OOptional(v A) Optional {
	return Optional{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Optional) Get(deflt A) A {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalTinyJSON does JSON marshaling using tinyjson interface.
func (v Optional) MarshalTinyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Optional(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalTinyJSON does JSON unmarshaling using tinyjson interface.
func (v *Optional) UnmarshalTinyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Optional{}
	} else {
		v.V = l.Optional()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Optional) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalTinyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Optional) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalTinyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Optional) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Optional) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
