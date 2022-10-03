package model_test

import (
	"testing"

	"github.com/w-woong/user/pkg/core/model"
)

func getEmptyUser() model.User {
	return model.NilUser
}

func TestUserIsNil(t *testing.T) {
	user := getEmptyUser()
	if !user.IsNil() {
		t.FailNow()
	}
}
