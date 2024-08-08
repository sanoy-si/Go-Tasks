package middleware

import (
	"fmt"
	"net/http"
	"time"

	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"Authorization header is required"})
			c.Abort()
			return
		}
		
		authParts := strings.Split(authHeader, " ")
		if len(authParts) == 1 || strings.ToLower(authParts[0]) != "bearer"{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"Invalid authorizaton header"})
			c.Abort()
			return
		}


		keyFunc := func(token *jwt.Token) (interface{}, error){
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok{
				return nil, fmt.Errorf("unexpected sigining method %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil 
		}

		token, err := jwt.Parse(authParts[1], keyFunc)
		if err != nil || !token.Valid{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"Invalid JWT"})
			c.Abort()
			return

		}


		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"Invalid JWT"})
			c.Abort()
			return
		}

		if claims["expires_at"].(float64) < float64(time.Now().Local().Unix()){
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"token expired"})
			c.Abort()
			return
		}

		c.Set("is_admin",claims["is_admin"])

		c.Next()
	}
}


func AdminMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		if c.Value("is_admin") != true{
			c.IndentedJSON(http.StatusForbidden, gin.H{"error":"Allowed for admins only"})
			c.Abort()
			return
		}

		c.Next()
	}

}