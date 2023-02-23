package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/smalake/kakebo-api/utils/logging"
)

func CreateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 168).Unix(), //有効期限1週間
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return "jwt create faild", err
	}
	return tokenString, nil
}
