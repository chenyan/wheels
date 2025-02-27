package middlewares

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	HeaderKey = "X-T"
)

var (
	Secret                         = []byte("rhizome-Xj3L.")
	DefaultExpiration              = time.Hour * 24 * 30
	DefaultMaxAge                  = 60 * 60 * 24 * 30
	Logger            *slog.Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
)

type JWTSessionClaims struct {
	UserID   uint64         `json:"user_id,omitempty"`
	Username string         `json:"username,omitempty"`
	Extra    map[string]any `json:"extra,omitempty"`
	jwt.RegisteredClaims
}

// ParseJWT parses a JWT token and returns the claims as a map.
func ParseJWT(token string, secret []byte) (*JWTSessionClaims, error) {
	t, err := jwt.ParseWithClaims(token, &JWTSessionClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if !t.Valid {
		return nil, fmt.Errorf("invalid jwt token: %v", err)
	}

	if m, ok := t.Claims.(*JWTSessionClaims); ok {
		return m, nil
	} else {
		return nil, err
	}
}

// GinJWTMiddleware is a Gin middleware that parses JWT tokens from the
func GinJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader(HeaderKey)
		if jwtToken != "" {
			claims, err := ParseJWT(jwtToken, Secret)
			if err != nil {
				log.Printf("ginjwt: parse jwt token error: %v", err)
			} else {
				c.Set("user_id", claims.UserID)
				log.Println("ginjwt: user_id,", claims.UserID, "expired", claims.ExpiresAt.Time)
			}
			// TODO 做服务端session校验
		}
		c.Next()
	}
}

// GenJWTToken generates a JWT token with a given user ID and extra data.
func GenJWTToken(userID uint64, username string, extra map[string]any) (string, error) {
	// Create a new JWTClaims struct.
	claims := JWTSessionClaims{
		UserID:   userID,
		Username: username,
		Extra:    extra,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DefaultExpiration)),
		},
	}
	// Create a new JWT token with the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// Sign the token with the secret.
	ss, err := token.SignedString(Secret)
	log.Println(ss, err)
	return ss, err
}
