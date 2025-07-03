package middleware

import (
	"fmt"
	"net/http"
	"p2final/helper"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Missing Bearer token", nil)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return helper.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid or expired token", err.Error())
			}

			c.Set("user", token)
			return next(c)
		}
	}
}
