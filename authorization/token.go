package authorization

import (
	"errors"
	"github.com/GMaikerYactayo/crud-api-goland/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateToken -
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		MapClaims: jwt.MapClaims{
			"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour)).Unix(),
			"Issuer":    "Maiker",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, VerifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("invalid token")
	}
	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("claims could not be obtained")
	}
	return *claim, nil
}

func VerifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
