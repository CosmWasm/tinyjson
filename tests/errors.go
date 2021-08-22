package tests

//tinyjson:json
type ErrorIntSlice []int

//tinyjson:json
type ErrorBoolSlice []bool

//tinyjson:json
type ErrorUintSlice []uint

//tinyjson:json
type ErrorStruct struct {
	Int      int    `json:"int"`
	String   string `json:"string"`
	Slice    []int  `json:"slice"`
	IntSlice []int  `json:"int_slice"`
}

type ErrorNestedStruct struct {
	ErrorStruct ErrorStruct `json:"error_struct"`
	Int         int         `json:"int"`
}

//tinyjson:json
type ErrorIntMap map[uint32]string
