package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenDuration = 24 * time.Hour

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService interface {
	Generate(userID uint) (string, error)
	Validate(token string) (*Claims, error)
}

type jwtService struct {
	secret []byte
}

func NewJWTService(secret string) JWTService {
	return &jwtService{secret: []byte(secret)}
}

func (s *jwtService) Generate(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// O keyfunc valida o algoritmo antes de retornar a chave,
// evitando o ataque "alg:none" onde o token chega sem assinatura.
func (s *jwtService) Validate(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metodo inesperado")
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token invalido")
	}
	return claims, nil
}
