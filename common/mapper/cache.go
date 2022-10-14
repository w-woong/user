package mapper

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	_mapper sync.Map
)

func StoreMapper(src, dest any) {
	key := PairKey(src, dest)
	srcType := reflect.TypeOf(src).Elem()
	destType := reflect.TypeOf(dest).Elem()
	sm := NewStructMapper(srcType, destType)

	_mapper.Store(key, sm)
}

func LoadMapper(src, dest any) (*StructMapper, error) {
	key := PairKey(src, dest)
	sm, ok := _mapper.Load(key)
	if !ok {
		return nil, errors.New("mapper not found")
	}

	return sm.(*StructMapper), nil
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

func unpackEFace(obj interface{}) *eface {
	return (*eface)(unsafe.Pointer(&obj))
}

func getKey(v any) uintptr {
	return uintptr(unpackEFace(v).rtype)
}

func getPairKey(src, dest uintptr) string {
	return fmt.Sprintf("%x_%x", src, dest)
}

func PairKey(src, dest interface{}) string {
	srcKey := getKey(src)
	destKey := getKey(dest)

	return getPairKey(srcKey, destKey)
}
