package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Função para hashear a senha do usuário

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// definindo quanto tempo o token vai durar

func jwtTTL() time.Duration {

	if s := os.Getenv("JWT_TTL"); s != "" {
		if m, err := time.ParseDuration(s + "m"); err == nil {
			return m
		}
	}
	return 15 * time.Minute

}

// criando o token JWT

func CreateJWT(userID, email string) (map[string]interface{}, error) {

	secret := os.Getenv("JWT_SECRET")
	issuer := os.Getenv("JWT_ISSUER")
	expire := time.Now().Add(jwtTTL()).Unix()

	if issuer == "" {
		issuer = "myapp"
	}

	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"iss":   issuer,
		"iat":   time.Now().Unix(),
		"exp":   expire,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error signing token:", err)
		return nil, err
	}

	return map[string]interface{}{
		"token":   tokenString,
		"expires": expire,
	}, nil

}

// verificando o token JWT

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		uid, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token subject"})
			return
		}

		c.Set("userID", uid)
		c.Next()

	}
}
