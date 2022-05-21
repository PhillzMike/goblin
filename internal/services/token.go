package services

import (
	"fmt"
	"time"

	"github.com/Zaida-3dO/goblin/config"
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/repositories"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type token struct {
	accessToken  string
	accessUUID   string
	atExpiresIn  int64
	refreshToken string
	refreshUUID  string
	rtExpiresIn  int64
	emailToken   string
	etExpiresIn  int64
}

func NewToken(accessToken, refreshToken, emailToken, accessUUID,
	refreshUUID string, atExpiresIn, rtExpiresIn, etExpiresIn int64) *token {
	return &token{
		accessToken,
		accessUUID,
		atExpiresIn,
		refreshToken,
		refreshUUID,
		rtExpiresIn,
		emailToken,
		etExpiresIn,
	}
}

type TokenService interface {
	GenerateTokenPair(uint) (*token, *errs.Err)
	GetUserFromAccessToken(string) (dtos.User, *errs.Err)
	GenerateEmailToken(string) (*string, *errs.Err)
	GetEmailFromToken(string) (string, *errs.Err)
}

type tokenService struct {
	userRepo     repositories.UserRepo
	userRepoMock repositories.UserRepoMock
}

func NewTokenService(mode string) TokenService {
	var ts TokenService = &tokenService{
		userRepo:     repositories.NewUserRepo(mode),
		userRepoMock: repositories.UserRepoMock{},
	}
	return ts
}

func (ts *tokenService) GenerateTokenPair(userId uint) (*token, *errs.Err) {
	var t token
	t.atExpiresIn = time.Now().Add(time.Hour * 24).Unix()
	t.rtExpiresIn = time.Now().Add(time.Hour * 24 * 180).Unix()

	t.accessUUID = uuid.NewV4().String()
	t.refreshUUID = uuid.NewV4().String()

	if err := t.generateAccessToken(userId); err != nil {
		return nil, errs.NewInternalServerErr(err.Error(), err)
	}
	if err := t.generateRefreshToken(userId); err != nil {
		return nil, errs.NewInternalServerErr(err.Error(), err)
	}

	return &t, nil
}

func (ts *tokenService) GetUserFromAccessToken(tokenStr string) (dtos.User, *errs.Err) {
	token, err := VerifyToken(tokenStr, config.Cfg.ATSecret)
	if err != nil {
		return dtos.User{}, err
	}

	var claims *jwt.MapClaims
	claims, err = validateToken(token)
	if err != nil {
		return dtos.User{}, err
	}

	userId, ok := (*claims)["user_id"]
	if !ok {
		return dtos.User{}, errs.NewUnauthorizedErr("invalid token", nil)
	}

	user := dtos.User{ID: uint(userId.(float64))}
	if err := ts.userRepo.GetUser(&user); err != nil {
		return dtos.User{}, err
	}

	return user, nil
}

func (ts *tokenService) GenerateEmailToken(email string) (*string, *errs.Err) {
	var t token
	t.etExpiresIn = time.Now().Add(time.Hour * 12).Unix()

	if err := t.generateEmailToken(email); err != nil {
		return nil, errs.NewInternalServerErr(err.Error(), err)
	}

	return &t.emailToken, nil
}

func (ts *tokenService) GetEmailFromToken(tokenStr string) (string, *errs.Err) {
	token, err := VerifyToken(tokenStr, config.Cfg.ETSecret)
	if err != nil {
		return "", err
	}

	var claims *jwt.MapClaims
	claims, err = validateToken(token)
	if err != nil {
		return "", err
	}

	email, ok := (*claims)["email"].(string)
	if !ok {
		return "", errs.NewBadRequestErr("invalid token", nil)
	}

	return email, nil
}

func (t *token) generateAccessToken(userId uint) error {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = t.accessUUID
	atClaims["user_id"] = userId
	atClaims["exp"] = t.atExpiresIn
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	var err error
	t.accessToken, err = at.SignedString([]byte(config.Cfg.ATSecret))
	if err != nil {
		return err
	}

	return nil
}

func (t *token) generateRefreshToken(userId uint) error {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = t.refreshUUID
	rtClaims["user_id"] = userId
	rtClaims["exp"] = t.rtExpiresIn
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	var err error
	t.refreshToken, err = rt.SignedString([]byte(config.Cfg.RTSecret))
	if err != nil {
		return err
	}

	return nil
}

func (t *token) generateEmailToken(email string) error {
	etClaims := jwt.MapClaims{}
	etClaims["email"] = email
	etClaims["exp"] = t.etExpiresIn
	et := jwt.NewWithClaims(jwt.SigningMethodHS256, etClaims)

	var err error
	t.emailToken, err = et.SignedString([]byte(config.Cfg.ETSecret))
	if err != nil {
		return err
	}

	return nil
}

func VerifyToken(tokenStr string, secret string) (*jwt.Token, *errs.Err) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errs.NewUnauthorizedErr("cannot parse auth token", err)
	}
	return token, nil
}

func validateToken(token *jwt.Token) (*jwt.MapClaims, *errs.Err) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errs.NewInternalServerErr("cannot validate token", nil)
	}
	return &claims, nil
}
