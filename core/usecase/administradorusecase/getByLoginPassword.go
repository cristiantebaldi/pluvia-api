package administradorusecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (usecase usecase) GetByLoginPassword(
	loginRequest *dto.AdministradorLoginRequestBody,
) (*domain.JwtAuthToken, error) {
	administrador, err := usecase.repository.GetByLoginPassword(loginRequest)
	if err != nil {
		return nil, err
	}

	fiveMinutes := time.Now().Add(time.Minute * 5).Unix()
	jwtToken := domain.NewJwtToken(*administrador, fiveMinutes)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)
	tokenString, err := token.SignedString([]byte(viper.GetString(`hash.bcrypt`)))
	if err != nil {
		return nil, err
	}

	auth, err := usecase.authRepository.GetByAdministradorID(administrador.ID)
	var hash string

	if err != nil {
		if err.Error() == "Token not created yet" {
			senha := []byte(viper.GetString(`hash.bcrypt`))
			hashedPassword, _ := bcrypt.GenerateFromPassword(senha, bcrypt.DefaultCost)
			hash = string(hashedPassword)

			auth := domain.Auth{
				Type:            "refreshToken",
				Hash:            hash,
				Token:           tokenString,
				AdministradorID: administrador.ID,
				Revoked:         false,
			}
			err := usecase.authRepository.Create(auth)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if auth != nil {
		if auth.Revoked {
			return nil, fmt.Errorf("your administrador has been revoked")
		}

		hash = auth.Hash
	}

	jwtAuthToken := domain.NewJwtAuthToken(tokenString, hash)

	return &jwtAuthToken, nil
}
