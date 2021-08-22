module github.com/CosmWasm/tinyjson/benchmark

go 1.12

require (
	github.com/json-iterator/go v1.1.7
	github.com/CosmWasm/tinyjson v0.0.0
	github.com/pquerna/ffjson v0.0.0-20190813045741-dac163c6c0a9
	github.com/ugorji/go/codec v1.1.7
	github.com/ugorji/go/codec/codecgen v1.1.7
)

replace github.com/CosmWasm/tinyjson => ../
