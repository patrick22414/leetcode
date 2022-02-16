package main

type I1 interface {
	Method() int
}

type I2 interface {
	Method() int

	Another() int
}

type My struct {
	field int
}

func (m My) Method() int {
	return m.field
}

func (m My) Another() int {
	return m.field + 1
}

func normalFunc(z I2) int {
	return z.Another()
}
