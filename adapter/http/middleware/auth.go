package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-viper/mapstructure/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/spf13/viper"
)

// VerifyAuthWithoutPermissions is a place function to control the session in middleware
func VerifyAuthWithoutPermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			bearToken := c.Request.Header.Get(
				"Authorization",
			) // bear token must be 2 params -- Bearer <token>

			if isAuth, access := verifyAccessToken(bearToken); isAuth {
				c.Request = setContextData(c.Request, &access)
				c.Next()
				return
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}


// VerifyAuthWithPermissions is a place function to control the session in middleware
func VerifyAuthWithPermissions(
) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			bearToken := c.Request.Header.Get(
				"Authorization",
			) // bear token must be 2 params -- Bearer <token>

			if isAuth, access := verifyAccessToken(bearToken); isAuth {
				c.Request = setContextData(c.Request, &access)

				if access.Enable == true {
					c.Next()
					return
				}

				
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func verifyAccessToken(bearToken string) (bool, domain.Admin) {
	admin := domain.Admin{}

	tokenBearer, err := ParseBearerToken(bearToken)

	if err != nil {
		return false, admin
	}

	token, err := jwt.Parse(*tokenBearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(`hash.bcrypt`)), nil //crie em uma variavel de ambiente
	})

	if err != nil {
		return false, admin
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims["admin"], &admin)
	} else {
		return false, admin
	}

	return true, admin
}
