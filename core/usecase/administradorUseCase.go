package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type administradorUseCase struct {
	repository     domain.AdministradorRepository
	authRepository domain.AuthRepository
}

func NewAdministradorUseCase(repository domain.AdministradorRepository, authRepository domain.AuthRepository) domain.AdministradorUseCase {
	return &administradorUseCase{
		repository:     repository,
		authRepository: authRepository,
	}
}

func (u *administradorUseCase) Create(administrador *dto.AdministradorRequestBody) (*domain.Administrador, error) {

	return u.repository.Create(administrador)
}

func (u *administradorUseCase) Fetch() (*[]domain.Administrador, error) {
	return u.repository.Fetch()
}

func (u *administradorUseCase) GetByLoginPassword(loginRequest *dto.AdministradorLoginRequestBody) (*domain.JwtAuthToken, error) {
	administrador, err := u.repository.GetByLoginPassword(loginRequest)
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

	auth, err := u.authRepository.GetByAdministradorID(administrador.ID)
	var hash string

	if err != nil {
		if err.Error() == "token not created yet" {
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
			err := u.authRepository.Create(auth)

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