package infrastructure

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
)

func GenerateToken(user domain.User) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.UserName,
		"is_admin":   user.IsAdmin,
		"expires_at": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}