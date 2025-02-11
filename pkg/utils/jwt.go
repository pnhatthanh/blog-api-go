package utils

import (
	"blogapi/config"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ACCESS_TOKEN_EXPIRATION_TIME, _  = time.ParseDuration(config.GetEnvOrDefault("ACCESS_TOKEN", "15"))
	REFRESH_TOKEN_EXPIRATION_TIME, _ = time.ParseDuration(config.GetEnvOrDefault("ACCESS_TOKEN", "60"))
)

func GenerateToken(userId string) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(ACCESS_TOKEN_EXPIRATION_TIME).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(REFRESH_TOKEN_EXPIRATION_TIME).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	secretKey := []byte(config.GetEnvOrDefault("JWT_SECRET_KEY", ""))
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
}

func ValidateRefershToken(token string)(jwt.MapClaims,error){
	parsedToken,err:=jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err!=nil{
		return nil, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		return nil, errors.New("invalid claim token")
	}
	return claims, nil
}

func GetUserIdByToken(tokenString string) (string, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}
	return userID, nil
}
