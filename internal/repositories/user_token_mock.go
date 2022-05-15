package repositories

import (
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type UserTokenRepoMock struct{}

var (
	userTokenDB map[uint]dtos.UserToken
	userTokenID uint
)

func (utr *UserTokenRepoMock) CreateToken(userToken dtos.UserToken) *errs.Err {
	userTokenDB = make(map[uint]dtos.UserToken)
	userToken.ID = userID
	userTokenDB[userToken.ID] = userToken
	userTokenID++
	return nil
}
