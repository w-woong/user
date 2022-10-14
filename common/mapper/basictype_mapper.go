package mapper

import (
	"reflect"
)

type BasicTypeMapper struct {
	SrcIndex  []int
	DestIndex []int
}

func (c *BasicTypeMapper) Map(srcValue, destValue reflect.Value) {
	srcField := srcValue.FieldByIndex(c.SrcIndex)
	destField := destValue.FieldByIndex(c.DestIndex)

	destField.Set(srcField)
}
