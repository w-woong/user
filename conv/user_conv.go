package conv

import (
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/user/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&commondto.User{}, &entity.User{})
	structmapper.StoreMapper(&entity.User{}, &commondto.User{})
}

// ToUserEntity converts dto.User to entity.User
func ToUserEntity(src *commondto.User) (user entity.User, err error) {
	// src.BirthDate = time.Date(src.BirthYear, time.Month(src.BirthMonth), src.BirthDay, 0, 0, 0, 0, time.UTC)

	err = structmapper.Map(src, &user)
	user.Personal.CombineBirthdate()
	return
}

// ToUserDto converts entity.User to dto.User.
func ToUserDto(src *entity.User) (user commondto.User, err error) {
	err = structmapper.Map(src, &user)
	return
}
