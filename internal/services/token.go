package services

import (
	"fmt"
	"time"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"github.com/Zaida-3dO/goblin/config"
	"github.com/twinj/uuid"
	"github.com/dgrijalva/jwt-go"
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
	GenerateTokenPair(userId uint) (*token, *errs.Err)
}

type tokenService struct{}

func NewTokenService() TokenService {
	var ts TokenService = &tokenService{}
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
