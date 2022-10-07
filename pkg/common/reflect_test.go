package common_test

import (
	"encoding/json"
	"fmt"
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

	common.ScanStruct(&src, &dst)
	assert.EqualValues(t, src, dst)
}

type MyStr string
type StructA struct {
	Str string
}

type StructB struct {
	Str MyStr
}

func TestScanStructWithTypeAlias(t *testing.T) {
	a := StructA{
		Str: "hello",
	}
	b := StructB{}

	common.ScanStruct(&a, &b)
	assert.EqualValues(t, a.Str, b.Str)

	c := StructB{"hello"}
	d := StructA{}
	common.ScanStruct(&c, &d)
	assert.EqualValues(t, c.Str, d.Str)
}

type CountryCode string

type UserDto struct {
	EmailsPtr   []*EmailDto
	ID          string
	Name        *string
	Nationality string
	Email       EmailDto
	EmailPtr    *EmailDto
	Emails      []EmailDto
}

type EmailDto struct {
	Addr string
}

type User struct {
	ID          string
	Name        *string
	Nationality CountryCode
	Email       Email
	EmailPtr    *Email
	Emails      []Email
	EmailsPtr   []*Email
}

func (e *User) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Email struct {
	Addr string
}

func TestScanStruct(t *testing.T) {

	ID := "wonksing"
	name := "wonk"
	emails := make([]EmailDto, 0)
	emails = append(emails, EmailDto{Addr: "wonk@wonk.orgg"})
	emailsPtr := make([]*EmailDto, 0)
	emailsPtr = append(emailsPtr, &EmailDto{Addr: "wonk@wonk.orgg"})

	userDto := UserDto{
		ID:          ID,
		Name:        &name,
		Nationality: "KOR",
		Email: EmailDto{
			Addr: "wonk@wonk.orgg",
		},
		EmailPtr: &EmailDto{
			Addr: "ptrwonk@wonk.orgg",
		},
		Emails:    emails,
		EmailsPtr: emailsPtr,
	}

	user := User{}
	common.ScanStruct(&userDto, &user)
	fmt.Println("reflected user:", user.String())
	assert.EqualValues(t, `{"ID":"wonksing","Name":"wonk","Nationality":"KOR","Email":{"Addr":"wonk@wonk.orgg"},"EmailPtr":{"Addr":"ptrwonk@wonk.orgg"},"Emails":[{"Addr":"wonk@wonk.orgg"}],"EmailsPtr":[{"Addr":"wonk@wonk.orgg"}]}`,
		user.String())
}

func BenchmarkScanStruct(b *testing.B) {

	ID := "wonksing"
	name := "wonk"
	emails := make([]EmailDto, 0)
	emails = append(emails, EmailDto{Addr: "wonk@wonk.orgg"})
	emailsPtr := make([]*EmailDto, 0)
	emailsPtr = append(emailsPtr, &EmailDto{Addr: "wonk@wonk.orgg"})

	userDto := UserDto{
		ID:          ID,
		Name:        &name,
		Nationality: "KOR",
		Email: EmailDto{
			Addr: "wonk@wonk.orgg",
		},
		EmailPtr: &EmailDto{
			Addr: "ptrwonk@wonk.orgg",
		},
		Emails:    emails,
		EmailsPtr: emailsPtr,
	}
	for i := 0; i < b.N; i++ {
		user := User{}
		common.ScanStruct(&userDto, &user)
	}

}
