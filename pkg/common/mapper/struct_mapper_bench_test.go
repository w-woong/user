package mapper_test

import (
	"reflect"
	"testing"

	"github.com/w-woong/user/pkg/common/mapper"
)

func BenchmarkMapper(b *testing.B) {
	src := reflect.TypeOf(Person{})
	dest := reflect.TypeOf(PersonEntity{})
	sm := mapper.NewStructMapper(src, dest)

	mobiles := make([]*Mobile, 0)
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})
	mobiles = append(mobiles, &Mobile{Number: "210232323232", Provider: &Provider{Name: "ds"}})

	person := Person{
		Name:      "wonk",
		MobilePtr: &Mobile{Number: "010"},
		Mobile:    Mobile{Number: "20202"},
		Mobiles:   mobiles,
	}
	personEntity := PersonEntity{}

	for i := 0; i < b.N; i++ {
		sm.Map(&person, &personEntity)
	}
}
