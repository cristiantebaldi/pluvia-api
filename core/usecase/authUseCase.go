package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/spf13/viper"
)

type authUseCase struct {
	repository domain.AuthRepository
}

func NewAuthUseCase(repository domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{
		repository: repository,
	}
}

func (u *authUseCase) RefreshToken(refreshToken string) (*domain.JwtAuthToken, error) {
	auth, err := u.repository.GetByRefreshTypeToken("refreshToken", refreshToken)
	if err != nil {
		return nil, err
	}

	fiveMinutes := time.Now().Add(time.Minute * 5).Unix()
	jwtToken := domain.NewJwtToken(domain.Administrador{ID: auth.AdministradorID}, fiveMinutes)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)

	tokenString, err := token.SignedString([]byte(viper.GetString(`hash.bcrypt`)))
	if err != nil {
		return nil, err
	}

	auth.Token = tokenString
	err = u.repository.Update(auth.ID, *auth)
	if err != nil {
		return nil, err
	}

	jwtAuthToken := domain.NewJwtAuthToken(tokenString, auth.Hash)

	return &jwtAuthToken, nil
}