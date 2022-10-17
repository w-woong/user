package conv

import (
	"github.com/w-woong/user/dto"
	"github.com/w-woong/user/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.User{}, &entity.User{})
	structmapper.StoreMapper(&entity.User{}, &dto.User{})
}

// ToUserEntity converts dto.User to entity.User
func ToUserEntity(src *dto.User) (user entity.User, err error) {
	// src.BirthDate = time.Date(src.BirthYear, time.Month(src.BirthMonth), src.BirthDay, 0, 0, 0, 0, time.UTC)

	err = structmapper.Map(src, &user)
	user.Personal.CombineBirthdate()
	return
}

// ToUserDto converts entity.User to dto.User.
func ToUserDto(src *entity.User) (user dto.User, err error) {
	err = structmapper.Map(src, &user)
	return
}
