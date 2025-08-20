package adminusecase

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/spf13/viper"
)

func (usecase usecase) RefreshToken(refreshToken string, typeToken string) (*domain.JwtAuthToken, error) {
	auth, err := usecase.authRepository.GetByRefreshTypeToken(refreshToken, typeToken)
	if err != nil {
		return nil, err
	}

	if auth.Revoked {
		return nil, fmt.Errorf("this token is revoked")
	}

	admin, err := usecase.repository.GetByID(auth.AdminID)
	if err != nil {
		return nil, err
	}

	fiveMinutes := time.Now().Add(time.Minute * 5).Unix()
	jwtToken := domain.NewJwtToken(*admin, fiveMinutes)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)
	tokenString, err := token.SignedString([]byte(viper.GetString(`hash.bcrypt`)))
	if err != nil {
		log.Fatalln(err)
	}

	jwtAuthToken := domain.NewJwtAuthToken(tokenString, auth.Hash)
	if err != nil {
		return nil, err
	}

	return &jwtAuthToken, nil
}
