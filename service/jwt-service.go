package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token,error)

}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims

}

type jwtService struct {
	secrteKey 	string
	issuer 		string
}

func (j jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		UserID:      "albert",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0,3,0).Unix(),
			Issuer: j.issuer,
			IssuedAt:time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t, err := token.SignedString([]byte(j.secrteKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		if _,ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,fmt.Errorf("Unexpected signing method %v", tkn.Header["alg"])
		}
		return []byte(j.secrteKey),nil
	})
}

func NewJWTService() JWTService {
	return &jwtService{
		secrteKey: getSecretKey(),
		issuer: "albertanugerah",
	}
}

func getSecretKey() string{
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != ""{
		secretKey = "albrt"
	}
	return secretKey
}