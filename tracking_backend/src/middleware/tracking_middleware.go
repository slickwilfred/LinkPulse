package middlware

import (
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"tracking_backend/src/tokens/jwt_token"
)

type Tracking_Middleware struct {
	JWTSecretKey    string
	AllowedOrigin   string
	AllowedReferrer string
}

var secretKey = []byte("thisneedstochane")
func NewMiddleware(secretKey, allowedOrigin, allowedReferrer string) *Tracking_Middleware {
	return &Tracking_Middleware{
		JWTSecretKey:    secretKey,
		AllowedOrigin:   allowedOrigin,
		AllowedReferrer: allowedReferrer,
	}
}

func (tm *Tracking_Middleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		cookie, err := r.Cookie("auth")
	}
}

func (tm *Tracking_Middleware) ValidateOriginAndReferrer (next http.Handler) http.Handler {

}

