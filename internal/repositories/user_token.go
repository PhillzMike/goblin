package repositories

import (
	"github.com/Zaida-3dO/goblin/adapters/driven/dbs"
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"gorm.io/gorm"
)

type UserTokenRepo interface {
	CreateToken(userToken dtos.UserToken) *errs.Err
}

type userTokenRepo struct {
	psql *gorm.DB
}

func NewUserTokenRepo(mode string) UserTokenRepo {
	if mode == "test" {
		return &UserTokenRepoMock{}
	}
	var utr UserTokenRepo = &userTokenRepo{
		psql: dbs.GetInstance(mode),
	}
	return utr
}

func (utr *userTokenRepo) CreateToken(userToken dtos.UserToken) *errs.Err {
	err := utr.psql.Create(&userToken).Error
	if err != nil {
		return errs.NewInternalServerErr(err.Error(), err)
	}
	return nil
}
