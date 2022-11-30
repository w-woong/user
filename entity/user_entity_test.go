package entity_test

import (
	"testing"

	"github.com/w-woong/user/entity"
)

func getEmptyUser() entity.User {
	return entity.NilUser
}

func TestUserIsNil(t *testing.T) {
	user := getEmptyUser()
	if !user.IsNil() {
		t.FailNow()
	}
}
