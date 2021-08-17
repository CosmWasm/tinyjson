package tests

//tinyjson:json
type EscStringStruct struct {
	A string `json:"a"`
}

//tinyjson:json
type EscIntStruct struct {
	A int `json:"a,string"`
}
