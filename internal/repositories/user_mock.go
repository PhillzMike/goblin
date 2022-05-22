package repositories

import (
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type UserRepoMock struct{}

var (
	userDB map[uint]dtos.User
	userID uint
)

func (urm *UserRepoMock) CreateUser(user dtos.User) *errs.Err {
	userDB = make(map[uint]dtos.User)
	user.ID = userID
	userDB[user.ID] = user
	userID++
	return nil
}

func (urm *UserRepoMock) GetUser(user *dtos.User) *errs.Err {
	usr, ok := userDB[user.ID]
	if !ok {
		return errs.NewNotFoundErr("user not found!", nil)
	}
	updateUser(user, usr)
	return nil
}

func (urm *UserRepoMock) FindUserByEmail(user *dtos.User, email string) *errs.Err {
	if userDB == nil {
		return errs.NewNotFoundErr("user not found!", nil)
	}

	for _, v := range userDB {
		if v.Email == email {
			updateUser(user, v)
			return nil
		}
	}
	return errs.NewNotFoundErr("user not found!", nil)
}

func (urm *UserRepoMock) SaveUser(user *dtos.User) *errs.Err {
	return nil
}

func (urm *UserRepoMock) ResetDB() {
	userDB = make(map[uint]dtos.User)
}

func updateUser(dest *dtos.User, src dtos.User) {
	dest.CreatedAt = src.CreatedAt
	dest.UpdatedAt = src.UpdatedAt
	dest.LastName = src.LastName
	dest.Email = src.Email
	dest.PhoneNumber = src.PhoneNumber
	dest.Password = src.Password
	dest.EmailVerificationCode = src.EmailVerificationCode
	dest.IsEmailVerified = src.IsEmailVerified
	dest.PasswordResetToken = src.PasswordResetToken
	dest.PasswordResetTokenExpiryDate = src.PasswordResetTokenExpiryDate
	dest.Colour = src.Colour
	dest.Gender = src.Gender
	dest.FirstName = src.FirstName
}
