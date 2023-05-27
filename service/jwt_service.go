package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issure:    "Chofongsua",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (s *jwtService) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	fmt.Printf("> secret: %v \n", s.secretKey)
	fmt.Printf("> claims: %v \n", *claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.secretKey))
	fmt.Printf("> %v \n", t)
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
