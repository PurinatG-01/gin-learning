package service

import (
	"fmt"
	model "gin-learning/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user model.User) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	DisplayName   string `json:"displayName"`
	DisplayImgUrl string `json:"displayImgUrl"`
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

func (s *jwtService) GenerateToken(user model.User) string {
	// Id:            user.Id,
	// 	Username:      user.Username,
	// 	DisplayName:   user.DisplayName,
	// 	DisplayImgUrl: user.DisplayImgUrl
	claims := &authCustomClaims{
		user.Id,
		user.Username,
		user.DisplayName,
		user.DisplayImgUrl,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
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
