package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(userId uint) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return "error", err
	}

	return t, nil
}

func ExtractJWT(c echo.Context) (jwt.MapClaims, error) {
	//Parse the token
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	// ... error handling
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func CheckAuthorization(id int, c echo.Context) bool {
	claims, _ := ExtractJWT(c)
	if int(claims["userId"].(float64)) == id {
		return true
	}
	return false
}
