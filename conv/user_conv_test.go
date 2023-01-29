package conv_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/user/conv"
)

func TestToUserEntity(t *testing.T) {
	password := commondto.CredentialPassword{
		Value: "asdfasdfasdfasdfasdfasdfasdf",
	}
	personal := commondto.Personal{
		FirstName:   "wonk",
		LastName:    "sun",
		BirthYear:   2022,
		BirthMonth:  1,
		BirthDay:    2,
		Gender:      "M",
		Nationality: "KOR",
	}
	emails := make([]commondto.Email, 0)
	emails = append(emails, commondto.Email{
		Email:    "wonk@wonk.orgg",
		Priority: 0,
	})
	emails = append(emails, commondto.Email{
		Email:    "monk@wonk.orgg",
		Priority: 1,
	})
	src := commondto.User{
		LoginID:            "wonk",
		LoginType:          "id",
		CredentialPassword: password,
		Personal:           personal,
		Emails:             emails,
	}

	res, err := conv.ToUserEntity(&src)
	assert.Nil(t, err)

	fmt.Println(res.String())
}

func TestFixedZone(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	utc := time.Now().In(loc)
	fmt.Println(utc)

	loc, _ = time.LoadLocation("Asia/Seoul")
	utc = time.Now().In(loc)
	fmt.Println(utc)

	loc, _ = time.LoadLocation("America/New_York")
	utc = time.Now().In(loc)
	fmt.Println(utc)
}

type Student struct {
	CreatedAt time.Time
}

func (s *Student) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
func unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func TestTimezone(t *testing.T) {
	s := Student{}
	v := "2022-10-15T09:00:00+00:00"
	tim, _ := time.Parse(time.RFC3339, v)
	fmt.Println(tim)

	jsonStr := `{"CreatedAt":"%s"}`
	unmarshal([]byte(fmt.Sprintf(jsonStr, v)), &s)
	fmt.Println(s.String())

	v = "2022-10-15T18:00:00+09:00"
	tim, _ = time.Parse(time.RFC3339, v)
	fmt.Println(tim)

	unmarshal([]byte(fmt.Sprintf(jsonStr, v)), &s)
	fmt.Println(s.String())

	v = "2022-10-15T05:00:00-04:00"
	tim, _ = time.Parse(time.RFC3339, v)
	fmt.Println(tim)

	unmarshal([]byte(fmt.Sprintf(jsonStr, v)), &s)
	fmt.Println(s.String())
}
