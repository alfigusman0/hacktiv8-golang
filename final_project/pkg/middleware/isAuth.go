package middleware

import (
	"errors"
	"final_project/pkg/models"
	"log"
	"net/http"
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
		errInvalidToken := errors.New("invalid token")
		if len(split) != 2 {
			ctx.AbortWithStatusJSON(401, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "error",
				"message": errInvalidToken.Error(),
			})
			return
		}
		getToken := split[1]
		var checkJwt models.Jwt
		err := db.Where("token = ? and expired = ?", getToken, "TIDAK").First(&checkJwt).Error
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"status":  "error",
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
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "error",
				"message": errInvalidToken.Error(),
			})
			return
		}
		if _, ok := validated.Claims.(jwt.MapClaims); !ok && !validated.Valid {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  "error",
				"message": errInvalidToken.Error(),
			})
			return
		}
		ctx.Set("user", validated.Claims.(jwt.MapClaims))
		ctx.Next()
	}
}
