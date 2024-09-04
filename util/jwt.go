package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/quest-be/constant"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	UserId    uint64    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func CreateToken(userId uint64, duration time.Duration) (string, error) {
	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(Default.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewPayload(userId uint64, duration time.Duration) (*Payload, error) {

	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		Id:        tokenId,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return constant.ErrExpiredToken
	}
	return nil
}
