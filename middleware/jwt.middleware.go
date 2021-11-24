package middleware

import (
	"context"
	"log"
	"net/http"
	"pustaka-api/JWT"
	"pustaka-api/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//* HELPER FUNC UNTUK DIPANGGIL LAGI NANTI, UNTUK STORE SEMENTARA
func addToContext(c *gin.Context, key interface{}, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService JWT.IJwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			res := helper.BuildErrorResponse("Token is not validssss", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			//* LANSUNG STORE USER ID NYA
			c.Request = addToContext(c, "GETUSER", claims["user_id"])
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["iss"])
			c.Next()
		}
	}
}

func GetCurrentUser(ctx context.Context) interface{} {
	cu := ctx.Value("GETUSER")
	return cu
}
