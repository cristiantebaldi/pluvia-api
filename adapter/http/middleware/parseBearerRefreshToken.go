package middleware

import (
	"fmt"
	"strings"
)

func ParseBearerRefreshToken(bearerToken string) (*string, *string, error) {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 4 {
		return nil, nil, fmt.Errorf("access denided")
	}

	refreshToken := strArr[2]
	typeToken := strArr[3]

	return &refreshToken, &typeToken, nil
}
