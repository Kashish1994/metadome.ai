package helper

import (
	"fmt"
	"github.com/eduhub/util"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	// Add other relevant claims as needed
}

func GenerateToken(email string) (string, error) {
	fmt.Println("email", email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
	fmt.Println("tokenClaims", token.Claims)
	var secretKey = []byte(util.SecretKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	fmt.Println("token", tokenString)
	return tokenString, nil
}

func DecodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(util.SecretKey), nil
	})

	if err != nil {
		fmt.Println("Error decoding token:", err)
		return nil, err
	}

	return token, nil
}
