package common

import (
	"reflect"
)

// ScanStruct scans src struct's fields and sets them to dest's fields that have the same field names.
func ScanStruct(src, dest interface{}) {
	srcValue := reflect.ValueOf(src).Elem()
	destValue := reflect.ValueOf(dest).Elem()

	scanStruct(srcValue, destValue)
}

func scan(kind reflect.Kind, src, dst reflect.Value) {
	switch kind {
	case reflect.Struct:
		scanStruct(src, dst)
	case reflect.Slice:
		scanSlice(src, dst)
	default:
		scanOthers(src, dst)
	}
}

func scanOthers(src reflect.Value, dst reflect.Value) {
	dstFieldType := dst.Type()
	if src.Type().AssignableTo(dstFieldType) {
		dst.Set(src)
	} else if src.CanConvert(dstFieldType) {
		dst.Set(src.Convert(dstFieldType))
	}
}

func scanSlice(src reflect.Value, dst reflect.Value) {
	dstFieldSlice := reflect.MakeSlice(reflect.SliceOf(dst.Type().Elem()), src.Len(), src.Cap())
	for j := 0; j < src.Len(); j++ {
		srcValue := reflect.Indirect(src.Index(j))
		dstValue := dstFieldSlice.Index(j)
		if dstValue.Kind() == reflect.Pointer {
			dstValue.Set(reflect.New(dstValue.Type().Elem()))
		}
		dstValue = reflect.Indirect(dstValue)
		scanStruct(srcValue, dstValue)
	}
	dst.Set(dstFieldSlice)
}

func scanStruct(src reflect.Value, dst reflect.Value) {
	n := src.NumField()
	for i := 0; i < n; i++ {
		f := src.Type().Field(i)
		if f.PkgPath != "" && f.Name != "_" {
			continue
		}

		srcFieldKind := f.Type.Kind()
		if srcFieldKind == reflect.Interface {
			continue
		}

		if _, ok := dst.Type().FieldByName(f.Name); !ok {
			continue
		}

		srcValue := src.Field(i)
		dstValue := dst.FieldByName(f.Name)

		scan(srcFieldKind, srcValue, dstValue)
	}
}
