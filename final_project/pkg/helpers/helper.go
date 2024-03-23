package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(hash), err

}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

type UsersClaims struct {
	Unik     int64  `json:"unik"`
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, nama string, username string) (string, error) {
	claims := UsersClaims{
		Unik:     time.Now().Unix(),
		ID:       id,
		Nama:     nama,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return signedToken, err
}
