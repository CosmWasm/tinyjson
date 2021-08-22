package tests

//tinyjson:json
type (
	GenDeclared1 struct {
		Value string
	}

	// A gen declared tinyjson struct with a comment
	GenDeclaredWithComment struct {
		Value string
	}
)

type (
	//tinyjson:json
	TypeDeclared struct {
		Value string
	}

	TypeNotDeclared struct {
		Value string
	}
)

var (
	myGenDeclaredValue             = TypeDeclared{Value: "GenDeclared"}
	myGenDeclaredString            = `{"Value":"GenDeclared"}`
	myGenDeclaredWithCommentValue  = TypeDeclared{Value: "GenDeclaredWithComment"}
	myGenDeclaredWithCommentString = `{"Value":"GenDeclaredWithComment"}`
	myTypeDeclaredValue            = TypeDeclared{Value: "TypeDeclared"}
	myTypeDeclaredString           = `{"Value":"TypeDeclared"}`
)
