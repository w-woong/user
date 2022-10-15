package dto_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/tj/assert"
	"github.com/w-woong/common"
	"github.com/w-woong/user/dto"
)

func TestCreateUser(t *testing.T) {
	userEmail := dto.UserEmail{
		Email: "wonk@wonk.orgg",
	}
	userEmails := []dto.UserEmail{userEmail}

	user := dto.User{
		LoginID:     "wonksing",
		FirstName:   "wonk",
		LastName:    "sun",
		BirthDate:   time.Now().Round(0),
		Gender:      "M",
		Nationality: "KOR",
		UserEmails:  userEmails,
	}
	reqBody := common.HttpBody{
		Document: &user,
	}
	b, err := json.Marshal(&reqBody)
	assert.Nil(t, err)

	fmt.Println(string(b))
}
