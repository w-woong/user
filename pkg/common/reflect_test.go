package common_test

import (
	"testing"

	"github.com/tj/assert"
	"github.com/w-woong/user/pkg/common"
)

type Src struct {
	InnerValue Inner
	Inner
	InnerSlice []Inner
	String     string
	Byte       byte
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Bool       bool

	Strings  []string
	Bytes    []byte
	Ints     []int
	Int8s    []int8
	Int16s   []int16
	Int32s   []int32
	Int64s   []int64
	Uints    []uint
	Uint8s   []uint8
	Uint16s  []uint16
	Uint32s  []uint32
	Uint64s  []uint64
	Float32s []float32
	Float64s []float64
	Bools    []bool

	StringPointer  *string
	BytePointer    *byte
	IntPointer     *int
	Int8Pointer    *int8
	Int16Pointer   *int16
	Int32Pointer   *int32
	Int64Pointer   *int64
	UintPointer    *uint
	Uint8Pointer   *uint8
	Uint16Pointer  *uint16
	Uint32Pointer  *uint32
	Uint64Pointer  *uint64
	Float32Pointer *float32
	Float64Pointer *float64
	BoolPointer    *bool

	StringsPointer  *[]string
	BytesPointer    *[]byte
	IntsPointer     *[]int
	Int8sPointer    *[]int8
	Int16sPointer   *[]int16
	Int32sPointer   *[]int32
	Int64sPointer   *[]int64
	UintsPointer    *[]uint
	Uint8sPointer   *[]uint8
	Uint16sPointer  *[]uint16
	Uint32sPointer  *[]uint32
	Uint64sPointer  *[]uint64
	Float32sPointer *[]float32
	Float64sPointer *[]float64
	BoolsPointer    *[]bool
}

type Dst struct {
	InnerValue Inner
	Inner
	InnerSlice []Inner
	String     string
	Byte       byte
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Bool       bool

	Strings  []string
	Bytes    []byte
	Ints     []int
	Int8s    []int8
	Int16s   []int16
	Int32s   []int32
	Int64s   []int64
	Uints    []uint
	Uint8s   []uint8
	Uint16s  []uint16
	Uint32s  []uint32
	Uint64s  []uint64
	Float32s []float32
	Float64s []float64
	Bools    []bool

	StringPointer  *string
	BytePointer    *byte
	IntPointer     *int
	Int8Pointer    *int8
	Int16Pointer   *int16
	Int32Pointer   *int32
	Int64Pointer   *int64
	UintPointer    *uint
	Uint8Pointer   *uint8
	Uint16Pointer  *uint16
	Uint32Pointer  *uint32
	Uint64Pointer  *uint64
	Float32Pointer *float32
	Float64Pointer *float64
	BoolPointer    *bool

	StringsPointer  *[]string
	BytesPointer    *[]byte
	IntsPointer     *[]int
	Int8sPointer    *[]int8
	Int16sPointer   *[]int16
	Int32sPointer   *[]int32
	Int64sPointer   *[]int64
	UintsPointer    *[]uint
	Uint8sPointer   *[]uint8
	Uint16sPointer  *[]uint16
	Uint32sPointer  *[]uint32
	Uint64sPointer  *[]uint64
	Float32sPointer *[]float32
	Float64sPointer *[]float64
	BoolsPointer    *[]bool
}

type Inner struct {
	String  string
	Byte    byte
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float32
	Float64 float64
	Bool    bool

	Strings  []string
	Bytes    []byte
	Ints     []int
	Int8s    []int8
	Int16s   []int16
	Int32s   []int32
	Int64s   []int64
	Uints    []uint
	Uint8s   []uint8
	Uint16s  []uint16
	Uint32s  []uint32
	Uint64s  []uint64
	Float32s []float32
	Float64s []float64
	Bools    []bool
}

func TestScan(t *testing.T) {
	str := "asdfasdf"
	strSlice := []string{"a", "s", "d", "f"}

	innerSlice := make([]Inner, 0)
	innerSlice = append(innerSlice, Inner{
		String: "aaabbb",
	})
	src := Src{
		String:         str,
		StringPointer:  &str,
		Strings:        strSlice,
		StringsPointer: &strSlice,
		InnerSlice:     innerSlice,
	}

	dst := Dst{}

	common.Scan(&src, &dst)
	assert.EqualValues(t, src, dst)
}
