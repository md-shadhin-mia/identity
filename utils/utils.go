package utils

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func Resources(r gin.RouterGroup, name string, c Controller) {
	r.GET("/"+name, c.GetAll)
	r.GET("/"+name+"/:id", c.GetById)
	r.POST("/"+name, c.Create)
	r.PUT("/"+name+"/:id", c.Update)
	r.DELETE("/"+name+"/:id", c.Delete)
}

func CopyCommonFields(src interface{}, dst interface{}) {
	// Get the reflect.Value and reflect.Type of the source and destination
	srcValue := reflect.ValueOf(src).Elem()
	srcType := srcValue.Type()

	dstValue := reflect.ValueOf(dst).Elem()
	// dstType := dstValue.Type()

	// Iterate over the fields of the source struct
	for i := 0; i < srcValue.NumField(); i++ {
		fieldName := srcType.Field(i).Name
		srcFieldValue := srcValue.Field(i)

		// Check if the destination struct has a field with the same name and type
		dstFieldValue := dstValue.FieldByName(fieldName)
		if dstFieldValue.IsValid() && dstFieldValue.CanSet() && dstFieldValue.Type() == srcFieldValue.Type() {
			dstFieldValue.Set(srcFieldValue)
		}
	}
}

func GenerateToken(userID uuid.UUID) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		bearer := strings.SplitN(tokenString, "Bearer ", 2)
		if len(bearer) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				log.Println("invalid token")
			}

			log.Println("user_id:", claims["user_id"])
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
