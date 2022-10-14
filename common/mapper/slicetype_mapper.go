package mapper

import "reflect"

type SliceTypeMapper struct {
	SrcIndex  []int
	DestIndex []int

	DestElementType   reflect.Type
	HasPointerElement bool
	ElementMapper     *StructMapper
}

func (c *SliceTypeMapper) Map(srcRootValue, destRootValue reflect.Value) {
	srcField := srcRootValue.FieldByIndex(c.SrcIndex)
	destField := destRootValue.FieldByIndex(c.DestIndex)

	dstFieldSlice := reflect.MakeSlice(c.DestElementType, srcField.Len(), srcField.Cap())
	for i := 0; i < srcField.Len(); i++ {
		srcElem := srcField.Index(i)
		destElem := dstFieldSlice.Index(i)
		if c.HasPointerElement {
			destElem.Set(reflect.New(destElem.Type().Elem()))
			c.ElementMapper.Map(srcElem.Interface(), destElem.Interface())
		} else {
			c.ElementMapper.Map(srcElem.Addr().Interface(), destElem.Addr().Interface())
		}
	}
	destField.Set(dstFieldSlice)
}
