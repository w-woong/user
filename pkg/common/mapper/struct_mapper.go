package mapper

import (
	"reflect"
)

func NewStructMapper(src, dest reflect.Type) *StructMapper {
	return NewStructMapperWithIndices([]int{}, []int{}, src, dest)
}

func NewStructMapperWithIndices(srcIndex, destIndex []int, src, dest reflect.Type) *StructMapper {
	sm := &StructMapper{
		SrcType:  src,
		DestType: dest,
	}
	sm.iterateStructFields(srcIndex, destIndex)

	return sm
}

type Mapper interface {
	Map(srcValue, destValue reflect.Value)
}

// StructMapper is to convert and set data between different kinds of structs.
// Each field is mapped by the field name, so they must be identical.
// Field type should be convertible if they are primary types.(eg. int8 and int, int and aliased int)
type StructMapper struct {
	// source type
	SrcType reflect.Type
	// destination type
	DestType reflect.Type

	// field indices to initialize source's fields
	srcIndicesToInitialize [][]int
	// field indices to initialize destination's fields
	destIndicesToInitialize [][]int

	// // mappers to convert basic type data
	// basicTypeMappers []BasicTypeMapper
	// // mappers to convert struct type data
	// structTypeMapper []StructTypeMapper
	// // mappers to convert slice type data
	// sliceTypeMapper []SliceTypeMapper

	mappers []Mapper
}

func (c *StructMapper) iterateStructFields(srcIndex, destIndex []int) {
	n := c.SrcType.NumField()
	for i := 0; i < n; i++ {
		srcField := c.SrcType.Field(i)
		if srcField.PkgPath != "" && srcField.Name != "_" {
			continue
		}

		srcFieldKind := srcField.Type.Kind()
		if srcFieldKind == reflect.Interface && srcField.Type.NumMethod() != 0 {
			continue
		}

		var destField reflect.StructField
		var ok bool
		if destField, ok = c.DestType.FieldByName(srcField.Name); !ok {
			continue
		}

		c.scanStructField(srcIndex, destIndex, srcFieldKind, srcField, destField)
	}
}

func (c *StructMapper) scanStructField(srcIndex, destIndex []int, srcFieldKind reflect.Kind, src, dest reflect.StructField) {
	switch srcFieldKind {
	case reflect.Pointer:
		switch src.Type.Elem().Kind() {
		case reflect.Struct:
			c.addStructPointerMapper(srcIndex, destIndex, src, dest)
		default:
			c.addBasicMapper(srcIndex, destIndex, src, dest)
		}
	case reflect.Struct:
		c.addStructMapper(srcIndex, destIndex, src, dest)
	case reflect.Slice:
		switch src.Type.Elem().Kind() {
		case reflect.Struct:
			c.addSliceOfStructMapper(srcIndex, destIndex, src, dest)
		case reflect.Pointer:
			if src.Type.Elem().Elem().Kind() == reflect.Struct {
				c.addSliceOfStructPointerMapper(srcIndex, destIndex, src, dest)
			} else {
				c.addBasicMapper(srcIndex, destIndex, src, dest)
			}
		default:
			c.addBasicMapper(srcIndex, destIndex, src, dest)
		}
	default:
		c.addBasicMapper(srcIndex, destIndex, src, dest)
	}
}

func (c *StructMapper) addStructPointerMapper(srcIndex, destIndex []int, src, dest reflect.StructField) {
	newSrcIndex := append(srcIndex, src.Index...)
	newDestIndex := append(destIndex, dest.Index...)
	c.mappers = append(c.mappers, &StructTypeMapper{
		SrcIndex:  newSrcIndex,
		DestIndex: newDestIndex,
		structMapper: NewStructMapperWithIndices(newSrcIndex, newDestIndex,
			src.Type.Elem(), dest.Type.Elem()),
	})

	c.srcIndicesToInitialize = append(c.srcIndicesToInitialize, newSrcIndex)
	c.destIndicesToInitialize = append(c.destIndicesToInitialize, newDestIndex)
}

func (c *StructMapper) addStructMapper(srcIndex, destIndex []int, src, dest reflect.StructField) {
	newSrcIndex := append(srcIndex, src.Index...)
	newDestIndex := append(destIndex, dest.Index...)
	c.mappers = append(c.mappers, &StructTypeMapper{
		SrcIndex:  newSrcIndex,
		DestIndex: newDestIndex,
		structMapper: NewStructMapperWithIndices(newSrcIndex, newDestIndex,
			src.Type, dest.Type),
	})
}

func (c *StructMapper) addBasicMapper(srcIndex, destIndex []int, src, dest reflect.StructField) {
	if src.Type.AssignableTo(dest.Type) || src.Type.ConvertibleTo(dest.Type) {
		fm := BasicTypeMapper{
			SrcIndex:  append(srcIndex, src.Index...),
			DestIndex: append(destIndex, dest.Index...),
		}
		c.mappers = append(c.mappers, &fm)
	}
}

func (c *StructMapper) addSliceOfStructMapper(srcIndex, destIndex []int, src, dest reflect.StructField) {
	c.mappers = append(c.mappers, &SliceTypeMapper{
		SrcIndex:          append(srcIndex, src.Index...),
		DestIndex:         append(destIndex, dest.Index...),
		DestElementType:   reflect.SliceOf(dest.Type.Elem()),
		HasPointerElement: dest.Type.Elem().Kind() == reflect.Pointer,
		ElementMapper:     NewStructMapper(src.Type.Elem(), dest.Type.Elem()),
	})
}

func (c *StructMapper) addSliceOfStructPointerMapper(srcIndex, destIndex []int, src, dest reflect.StructField) {
	c.mappers = append(c.mappers, &SliceTypeMapper{
		SrcIndex:          append(srcIndex, src.Index...),
		DestIndex:         append(destIndex, dest.Index...),
		DestElementType:   reflect.SliceOf(dest.Type.Elem()),
		HasPointerElement: dest.Type.Elem().Kind() == reflect.Pointer,
		ElementMapper:     NewStructMapper(src.Type.Elem().Elem(), dest.Type.Elem().Elem()),
	})
}

func (c *StructMapper) Map(src, dest interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	destValue := reflect.ValueOf(dest).Elem()

	initializeFieldsWithIndices(srcValue, c.srcIndicesToInitialize)
	initializeFieldsWithIndices(destValue, c.destIndicesToInitialize)

	for _, m := range c.mappers {
		m.Map(srcValue, destValue)
	}

}

// initializeFieldsWithIndices initializes directly using `indices`
func initializeFieldsWithIndices(v reflect.Value, indices [][]int) {
	for _, s := range indices {
		field := v.FieldByIndex(s)
		if field.Kind() == reflect.Pointer && field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
	}
}

// initializeStructFields traverses all fields recursivley and initialize nil pointer struct
func initializeStructFields(v reflect.Value) {
	if v.Type().Kind() != reflect.Struct {
		return
	}

	n := v.NumField()
	for i := 0; i < n; i++ {
		structField := v.Type().Field(i)
		// if !structField.IsExported() {
		// 	continue
		// }
		if structField.PkgPath != "" && structField.Name != "_" {
			continue
		}

		field := v.Field(i)

		// check if a field is pointer
		var fieldTypeKind reflect.Kind
		if field.Kind() == reflect.Pointer {
			fieldTypeKind = field.Type().Elem().Kind()
		} else {
			fieldTypeKind = field.Type().Kind()
		}

		// skip if it is not a struct
		if fieldTypeKind != reflect.Struct {
			continue
		}

		// initialize a field if it is nil
		var fieldValue reflect.Value
		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			fieldValue = field.Elem()
		} else {
			fieldValue = field
		}

		// initialize struct's nested field
		initializeStructFields(fieldValue)
	}
}
