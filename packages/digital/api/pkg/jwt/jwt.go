package jwt

import (
	"errors"
	"fmt"
	"github.com/cyberpunkattack/environment"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type UserClaim struct {
	Email 	  string `json:"email"`
	UserHash  string `json:"user_hash,omitempty"`
	Id        int    `json:"id,omitempty"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

const (
	RefreshTokenType = "refresh"
	AccessTokenType  = "access"
)

func CreateToken(email, userHash string, id int, tokenType string, expirationTime *time.Time) (string, error) {
	serverKey := environment.GEnv().Get("SERVER_KEY")

	claims := &UserClaim{
		Email: email,
		UserHash:         userHash,
		Id:               id,
		TokenType:        tokenType,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	if expirationTime != nil {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(*expirationTime)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(serverKey))
	fmt.Println(tokenString)
	return tokenString, err
}

func VerifyToken(tokenString string) (*UserClaim, error) {
	serverKey := environment.GEnv().Get("SERVER_KEY")
	claimsInstance := &UserClaim{}

	tToken, err := jwt.ParseWithClaims(tokenString, claimsInstance, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(serverKey), nil
	})

	if err != nil || !tToken.Valid {
		return nil, err
	}

	if !tToken.Valid {
		return nil, err
	}
	return claimsInstance, nil
}

func ValidateToken(tokenString string) bool {
	serverKey := environment.GEnv().Get("SERVER_KEY")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(serverKey), nil
	})

	if err != nil {
		return false
	}

	return true
}

func CheckIsTokenExpired(token string) bool {
	serverKey := environment.GEnv().Get("SERVER_KEY")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(serverKey), nil
	})

	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "token is expired")
}
