package dto_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tj/assert"
	"github.com/w-woong/user/pkg/dto"
)

func TestHttpBodyRegisterUser(t *testing.T) {
	user := dto.User{
		LoginID:     "wonksing",
		FirstName:   "wonk",
		LastName:    "sun",
		BirthDate:   "20000101",
		Gender:      "M",
		Nationality: "KOR",
	}
	reqBody := dto.HttpBody{
		Document: &user,
	}
	b, err := json.Marshal(&reqBody)
	assert.Nil(t, err)

	fmt.Println(string(b))
}
