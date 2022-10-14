package entity_test

import (
	"testing"

	"github.com/w-woong/user/dto"
)

func getEmptyUser() dto.User {
	return dto.NilUser
}

func TestUserIsNil(t *testing.T) {
	user := getEmptyUser()
	if !user.IsNil() {
		t.FailNow()
	}
}
