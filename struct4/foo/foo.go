package foo

type T struct {
	Field1 int
	Field2 int
}

func (t *T) Method1() int {
	return t.Field1
}

func (t *T) Method2() int {
	return t.Field2
}
