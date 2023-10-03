package str2type

func ResolveTypeSet(str ...string) map[GO_TYPE]struct{} {
	set := typeSet
	if !AssertBool(str...) {
		delete(set, BOOL)
	}
	if !AssertInt(str...) {
		delete(set, INT)
	}
	if !AssertFloat(str...) {
		delete(set, FLOAT)
	}
	return set
}

func ResolveTypeSingle(str ...string) GO_TYPE {
	switch {
	case AssertBool(str...):
		return BOOL
	case AssertInt(str...):
		return INT
	case AssertFloat(str...):
		return FLOAT
	}
	return STRING
}
