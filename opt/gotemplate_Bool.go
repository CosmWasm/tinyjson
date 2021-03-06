// generated by gotemplate

package opt

import (
	"fmt"

	"github.com/CosmWasm/tinyjson/jlexer"
	"github.com/CosmWasm/tinyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Bool struct {
	V       bool
	Defined bool
}

// Creates an optional type with a given value.
func OBool(v bool) Bool {
	return Bool{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Bool) Get(deflt bool) bool {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalTinyJSON does JSON marshaling using tinyjson interface.
func (v Bool) MarshalTinyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Bool(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalTinyJSON does JSON unmarshaling using tinyjson interface.
func (v *Bool) UnmarshalTinyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Bool{}
	} else {
		v.V = l.Bool()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Bool) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalTinyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Bool) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalTinyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Bool) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Bool) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
