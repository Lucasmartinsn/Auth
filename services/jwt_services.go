package services

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtservices struct {
	secretkey string
	issure    string
}

func NewJWTService() *jwtservices {
	return &jwtservices{
		secretkey: os.Getenv(""),
		issure:    "lol-api",
	}
}

type Clain struct {
	Sum int `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtservices) GenerateToken(id int) (string, error) {
	clain := &Clain{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)

	t, err := token.SignedString([]byte(s.secretkey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtservices) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretkey), nil
	})

	return err == nil
}

func (s *jwtservices) CreateJWT(data map[string]interface{}) (string, error) {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente para garantir que os números sejam realmente aleatórios
	randomNumber := rand.Intn(101)   // Gera um número inteiro aleatório entre 0 e 100
	// Configuração do token
	clain := &Clain{
		randomNumber,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)

	claims := token.Claims.(jwt.MapClaims)

	// Adição dos dados ao token
	for key, value := range data {
		claims[key] = value
	}

	// Definindo a expiração do token
	claims["exp"] = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    s.issure,
		IssuedAt:  time.Now().Unix(),
	}

	// Geração do token assinado
	tokenString, err := token.SignedString(s.secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}