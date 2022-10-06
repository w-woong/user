package common

import (
	"fmt"
	"reflect"
)

func ScanStruct(src interface{}, dst interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	srcValueType := srcValue.Type()
	dstValue := reflect.ValueOf(dst).Elem()
	dstValueType := dstValue.Type()

	n := srcValue.NumField()
	for i := 0; i < n; i++ {
		f := srcValueType.Field(i)
		if f.PkgPath != "" && f.Name != "_" {
			continue
		}

		k := f.Type.Kind()
		fmt.Println(k, f)
		if k == reflect.Interface {
			continue
		}
		// if k == reflect.Slice {
		// 	elemType := reflect.ValueOf(f.Type.Elem())
		// 	fmt.Println(elemType.Kind(), elemType)
		// }

		var dstField reflect.StructField
		var ok bool
		if dstField, ok = dstValueType.FieldByName(f.Name); !ok {
			continue
		}

		dstFieldType := dstField.Type
		if dstFieldType.Kind() != k {
			continue
		}

		srcField := srcValue.Field(i)
		if srcField.Type().AssignableTo(dstFieldType) {
			dstValue.FieldByName(f.Name).Set(srcField)
		} else if srcField.CanConvert(dstFieldType) {
			dstValue.FieldByName(f.Name).Set(srcField.Convert(dstFieldType))
		}
	}
}
