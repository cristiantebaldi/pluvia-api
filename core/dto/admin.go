package dto

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/util"
)

type AdminLoginRequestBody struct {
	Email    string `json:"email" example:"email" validate:"required"`
	Password string `json:"password" example:"the password" validate:"required,gte=5,lte=20"`
}

type AdminRequestBody struct {
	Name        string    `json:"name" validate:"required,gte=5,lte=150"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone" validate:"required,e164"`
	Enable      bool      `json:"enable" validate:"required"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func FromJsonToAdminLoginRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*AdminLoginRequestBody, error) {
	admin := AdminLoginRequestBody{}

	if err := json.NewDecoder(body).Decode(&admin); err != nil {
		return nil, err
	}
	if err := validate.Struct(admin); err != nil {
		return nil, util.HandleValidatorFieldError(err)
	}

	return &admin, nil
}

func FromJsonToAdminRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*AdminRequestBody, error) {
	admin := AdminRequestBody{}
	
	if err := json.NewDecoder(body).Decode(&admin); err != nil {
		return nil, err
	}

	if err := validate.Struct(admin); err != nil {
		return nil, util.HandleValidatorFieldError(err)
	}

	return &admin, nil
}