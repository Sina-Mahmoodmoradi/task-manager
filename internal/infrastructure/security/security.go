package security

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/core/port"
)


type JWTTokenManager struct{
	secret []byte
}

func NewJWTTokenManager() port.TokenManager {
	return &JWTTokenManager{
		secret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (m *JWTTokenManager) CreateToken(userID uint, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.secret)
}

func (m *JWTTokenManager) ParseToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return m.secret, nil
	})

	if err != nil {
		return 0, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return uint(userID), nil
		}
	}

	return 0, jwt.ErrTokenMalformed
}
