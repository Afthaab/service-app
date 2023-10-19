package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	privatekey *rsa.PrivateKey
	publickey  *rsa.PublicKey
}

type ctxKey int

const Key ctxKey = 1

func NewAuth(privatekey *rsa.PrivateKey, publickey *rsa.PublicKey) (*Auth, error) {
	if privatekey == nil || publickey == nil {
		return nil, errors.New("private/public key cannot be nil")
	}
	return &Auth{
		privatekey: privatekey,
		publickey:  publickey,
	}, nil
}

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenstr, err := tkn.SignedString(a.privatekey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}
	return tokenstr, nil
}

func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	tkn, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return a.publickey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}
	if !tkn.Valid {
		return jwt.RegisteredClaims{}, fmt.Errorf("invalid token")
	}
	return c, nil

}
