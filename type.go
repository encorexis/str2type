package str2type

type GO_TYPE int

const (
	BOOL GO_TYPE = iota
	INT
	FLOAT
	STRING
)

var typeSet = map[GO_TYPE]struct{}{
	BOOL:   {},
	INT:    {},
	FLOAT:  {},
	STRING: {},
}
