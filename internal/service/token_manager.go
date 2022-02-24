package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	NewJWT(userID string) (string, error)
	Parse(accessToken string) (*int64, error)
	NewRefreshToken() (string, error)
}

type tokenManager struct {
	signingKey string
}

func NewTokenManager(signedKey string) TokenManager {
	return &tokenManager{signingKey: signedKey}
}

func (t *tokenManager) NewJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60 * 24 * 30)),
		Subject:   userID,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *tokenManager) Parse(accessToken string) (*int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(t.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot get claims from token")
	}

	atoi, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("cannot convert str to int: %v", err)
	}

	id := int64(atoi)
	return &id, nil
}

func (t *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
