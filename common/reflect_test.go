package common_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/tj/assert"
	"github.com/w-woong/user/common"
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
	UpdatedAt    time.Time
	UpdatedAtPtr *time.Time
	CreatedAt    sql.NullTime
	CreatedAtPtr *sql.NullTime
	EmailsPtr    []*EmailDto
	ID           string
	Name         *string
	Nationality  string
	Email        EmailDto
	EmailPtr     *EmailDto
	Emails       []EmailDto
}

func (e *UserDto) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type EmailDto struct {
	Addr string
}

type User struct {
	ID           string
	Name         *string
	Nationality  CountryCode
	Email        Email
	EmailPtr     *Email
	Emails       []Email
	EmailsPtr    []*Email
	CreatedAt    sql.NullTime
	CreatedAtPtr *sql.NullTime
	UpdatedAt    time.Time
	UpdatedAtPtr *time.Time
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
	now, _ := time.Parse("20060102", "20221001")
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
		CreatedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
		CreatedAtPtr: &sql.NullTime{
			Time:  now,
			Valid: true},
		UpdatedAt:    now,
		UpdatedAtPtr: &now,
	}

	user := User{}
	common.ScanStruct(&userDto, &user)
	// fmt.Println("original user:", userDto.String())
	// fmt.Println("reflected user:", user.String())

	expected := `{"ID":"wonksing","Name":"wonk","Nationality":"KOR","Email":{"Addr":"wonk@wonk.orgg"},"EmailPtr":{"Addr":"ptrwonk@wonk.orgg"},"Emails":[{"Addr":"wonk@wonk.orgg"}],"EmailsPtr":[{"Addr":"wonk@wonk.orgg"}],"CreatedAt":{"Time":"2022-10-01T00:00:00Z","Valid":true},"CreatedAtPtr":{"Time":"2022-10-01T00:00:00Z","Valid":true},"UpdatedAt":"2022-10-01T00:00:00Z","UpdatedAtPtr":"2022-10-01T00:00:00Z"}`
	assert.EqualValues(t, expected, user.String())
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

type House struct {
	People People
}

func (d *House) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

type HouseEntity struct {
	People PeopleEntity
}

func (e *HouseEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type Person struct {
	Name string
}
type People []Person

type PersonEntity struct {
	Name string
}
type PeopleEntity []PersonEntity

func TestSliceTypeAlias(t *testing.T) {
	people := make(People, 0)
	people = append(people, Person{Name: "wonk"})
	house := House{People: people}

	var houseEntity HouseEntity
	common.ScanStruct(&house, &houseEntity)
	assert.Equal(t, house.String(), houseEntity.String())

	peopleEntity := make(PeopleEntity, 0)
	peopleEntity = append(peopleEntity, PersonEntity{Name: "mink"})
	houseEntity = HouseEntity{People: peopleEntity}
	house = House{}
	common.ScanStruct(&houseEntity, &house)
	assert.Equal(t, houseEntity.String(), house.String())
}
