package utilities

import (
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) domain.User {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return domain.User{}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return domain.User{}
	}

	issuer, err := claims.GetIssuer()
	if err != nil {
		return domain.User{}
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return domain.User{}
	}

	return domain.User{
		Email: issuer,
		ID:    subject,
	}
}
