package helper

import (
	"fmt"
	"pustaka-api/JWT"
	"github.com/dgrijalva/jwt-go"
)

//* FUNSI DIBAWAH TIDKA DIPAKAI KARENA SUDAH PAKAI AMBIL USER ID LANSUNG DI MIDDLEWARE
func GetUserIDByToken(token string, jwtServ JWT.IJwtService) (string, error) {
	aToken, err := jwtServ.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id, err
}
