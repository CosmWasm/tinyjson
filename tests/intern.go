package tests

//tinyjson:json
type NoIntern struct {
	Field string `json:"field"`
}

//tinyjson:json
type Intern struct {
	Field string `json:"field,intern"`
}

var intern = Intern{Field: "interned"}
var internString = `{"field":"interned"}`
