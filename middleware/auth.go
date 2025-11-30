package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	// if secret == "" {
	// 	secret = "your-secret-key-change-in-production" // Default for development
	// }
	jwtSecret = []byte(secret)
}

// Claims struct for JWT
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// AuthMiddleware checks for valid JWT token in cookie
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Validate token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user info in context for use in handlers
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// SetTokenCookie sets the JWT token as an HTTP-only cookie
func SetTokenCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"token",          // name
		token,            // value
		60*60*24,         // maxAge (24 hours in seconds)
		"/",              // path
		"",               // domain (empty = current domain)
		false,            // secure (set true in production with HTTPS)
		true,             // httpOnly (prevents JavaScript access)
	)
}

// ClearTokenCookie removes the JWT cookie (for logout)
func ClearTokenCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"token",
		"",
		-1,    // maxAge -1 deletes the cookie
		"/",
		"",
		false,
		true,
	)
}
