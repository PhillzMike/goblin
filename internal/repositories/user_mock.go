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

func (urm *UserRepoMock) ResetDB() {
	userDB = make(map[uint]dtos.User)
}

func updateUser(user1 *dtos.User, user2 dtos.User) {
	user1.CreatedAt = user2.CreatedAt
	user1.UpdatedAt = user2.UpdatedAt
	user1.LastName = user2.LastName
	user1.Email = user2.Email
	user1.PhoneNumber = user2.PhoneNumber
	user1.Password = user2.Password
	user1.EmailVerificationCode = user2.EmailVerificationCode
	user1.IsEmailVerified = user2.IsEmailVerified
	user1.PasswordResetToken = user2.PasswordResetToken
	user1.PasswordResetTokenExpiryDate = user2.PasswordResetTokenExpiryDate
	user1.Colour = user2.Colour
	user1.Gender = user2.Gender
	user1.FirstName = user2.FirstName
}
