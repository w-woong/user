package dto_test

import (
	"fmt"
	"reflect"
)

func scan(src interface{}, dst interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	srcType := reflect.TypeOf(src).Elem()
	// dstType := reflect.TypeOf(dst).Elem()

	n := srcType.NumField()
	for i := 0; i < n; i++ {
		f := srcType.Field(i)
		if f.PkgPath != "" && f.Name != "_" {
			continue
		}

		fmt.Println(f, reflect.Indirect(reflect.ValueOf(f)).Interface(), reflect.ValueOf(f).Interface())
		fmt.Println(srcValue.Field(i).Interface())
		// if dstField, ok := dstType.FieldByName(f.Name); ok {
		// 	f.
		// }
	}
}

func scan2(src interface{}, dst interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	// srcType := reflect.TypeOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()

	n := srcValue.NumField()
	for i := 0; i < n; i++ {
		f := srcValue.Type().Field(i)
		if f.PkgPath != "" && f.Name != "_" {
			continue
		}

		k := f.Type.Kind()
		if k == reflect.Interface {
			continue
		}

		if dstField, ok := dstValue.Type().FieldByName(f.Name); !ok {
			continue
		} else if dstField.Type.Kind() != k {
			continue
		}

		// fmt.Println(f, srcValue.Field(i).Interface())
		dstValue.FieldByName(f.Name).Set(srcValue.Field(i))
	}
}
