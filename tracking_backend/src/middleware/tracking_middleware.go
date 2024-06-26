package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Tracking_Middleware struct {
	JWTSecretKey    []byte
	AllowedOrigin   string
	AllowedReferrer string
	jwt.RegisteredClaims
}

func NewMiddleware(secretKey, allowedOrigin, allowedReferrer string) *Tracking_Middleware {
	return &Tracking_Middleware{
		JWTSecretKey:    []byte(secretKey),
		AllowedOrigin:   allowedOrigin,
		AllowedReferrer: allowedReferrer,
	}
}

func (tm *Tracking_Middleware) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"error": "No token provided"})
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return tm.JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("claims", token.Claims)
		c.Next()
	}
}

func (tm *Tracking_Middleware) ValidateOriginAndReferrer() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		referer := c.Request.Header.Get("Referer")

		if origin != tm.AllowedOrigin || !strings.HasPrefix(referer, tm.AllowedReferrer) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
