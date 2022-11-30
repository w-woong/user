package dto_test

import (
	"testing"
	"time"

	"github.com/w-woong/user/dto"
)

var (
	userDto = dto.User{
		ID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
		LoginID: "wonk",
		Password: dto.Password{
			ID:     "333cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID: "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			Value:  "asdfasdfasdf",
		},
		Personal: dto.Personal{
			ID:          "433cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			FirstName:   "wonk",
			LastName:    "sun",
			BirthYear:   2002,
			BirthMonth:  1,
			BirthDay:    2,
			BirthDate:   time.Date(2002, 1, 2, 0, 0, 0, 0, time.Local),
			Gender:      "M",
			Nationality: "KOR",
		},
	}
)

func TestUserString(t *testing.T) {
	// expected := `{"id":"22bcbf79-ca5f-42dc-8ca0-29441209a36a","login_id":"wonk","password":{"id":"333cbf79-ca5f-42dc-8ca0-29441209a36a","user_id":"22bcbf79-ca5f-42dc-8ca0-29441209a36a","value":"asdfasdfasdf"},"personal":{"id":"433cbf79-ca5f-42dc-8ca0-29441209a36a","user_id":"22bcbf79-ca5f-42dc-8ca0-29441209a36a","first_name":"wonk","last_name":"sun","birth_year":2002,"birth_month":1,"birth_day":2,"birth_date":"2002-01-02T00:00:00+09:00","gender":"M","nationality":"KOR"}}`
	// fmt.Println(userDto.String())
	// assert.EqualValues(t, expected, userDto.String())
}
