package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kadr/globe_express/config"
)

type JWT struct {
	appConfig *config.AppConfig
}

func NewJWT(appConfig *config.AppConfig) JWT {
	return JWT{appConfig: appConfig}
}

func (j *JWT) checkValid(tokenString string) error {
	token, err := j.prepareToken(tokenString)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return errors.New("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return errors.New("Invalid token signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			return errors.New("Token expired")
		default:
			return errors.New("Forbidden")
		}
	}
	if !token.Valid {
		return errors.New("Invalid token")
	}
	return nil
}

func (j *JWT) getUserID(tokenString string) (uuid.UUID, error) {
	token, err := j.prepareToken(tokenString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Can't parse token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userID, ok := claims["user_id"]; ok {
			if id, err := uuid.Parse(userID.(string)); err == nil {
				return id, nil
			} else {
				return uuid.Nil, fmt.Errorf("Incorrect user_id")
			}
		} else {
			return uuid.Nil, fmt.Errorf("user_id not found in token")
		}
	} else {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}
}

func (j *JWT) getUserRole(tokenString string) string {
	token, err := j.prepareToken(tokenString)
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userRole, ok := claims["user_role"]; ok {
				return userRole.(string)
			}
		}
	}
	return ""
}

func (j *JWT) prepareToken(tokenString string) (*jwt.Token, error) {
	if strings.Contains(tokenString, "Bearer") {
		tokenString = strings.Replace(tokenString, "Bearer", "", 1)
	}
	tokenString = strings.TrimSpace(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(j.appConfig.JWTSecret), nil
	})

	return token, err
}
