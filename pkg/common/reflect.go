package common

import "reflect"

func Scan(src interface{}, dst interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
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
		dstValue.FieldByName(f.Name).Set(srcValue.Field(i))
	}
}
