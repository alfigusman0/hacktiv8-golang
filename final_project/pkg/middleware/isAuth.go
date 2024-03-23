package middleware

import (
	"errors"
	"final_project/pkg/models"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func IsAuth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getHeader := ctx.GetHeader("Authorization")
		split := strings.Split(getHeader, "Bearer ")
		log.Println(split)
		errInvalidToken := errors.New("invalid token")
		if len(split) != 2 {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})
			return
		}
		getToken := split[1]
		var checkJwt models.Jwt
		err := db.Where("token = ? and expired = ?", getToken, "TIDAK").First(&checkJwt).Error
		log.Println(err)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})
			return
		}
		validated, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errInvalidToken
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message verify": errInvalidToken.Error(),
			})
			return
		}
		if _, ok := validated.Claims.(jwt.MapClaims); !ok && !validated.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})
			return
		}
		ctx.Set("user", validated.Claims.(jwt.MapClaims))
		ctx.Next()
	}
}
