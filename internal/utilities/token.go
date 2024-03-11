package utilities

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateToken(user domain.User) string {
	claims := CustomClaims{
		jwt.RegisteredClaims{
			Subject:   user.ID,
			Issuer:    user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("and0LXRlc3QK"))
	if err != nil {
		fmt.Println("err: ", err)
		signedToken = ""
	}

	return signedToken
}
