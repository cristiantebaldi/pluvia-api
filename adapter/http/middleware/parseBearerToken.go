package middleware

import (
	"fmt"
	"strings"
)

func ParseBearerToken(bearerToken string) (*string, error) {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 {
		return nil, fmt.Errorf("access denided")
	}

	token := strArr[1]

	return &token, nil
}
