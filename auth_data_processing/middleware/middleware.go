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
func JWTAuthentication(publicRoutes map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.FullPath()
		routeKey := fmt.Sprintf("%s %s", method, path)

		// verificando se a rota e publica

		if publicRoutes[routeKey] {
			c.Next()
			return
		}

		// exigindo autenticacao se nao for publica

		authHeader := c.GetHeader("Authorization")
		log.Println("Teste")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token nao fornecido"})
			return
		}

		// pegando o token do header
		tokenString := strings.Split(authHeader, " ")[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Verifica se o token expirou
		exp, ok := claims["exp"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		expirationTime := time.Unix(int64(exp), 0)
		if expirationTime.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			c.Abort()
			return
		}

		// pegando o id do ususario no token para colocar no contexto

		var userID string
		if sub, ok := claims["sub"].(string); ok {
			userID = sub
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		log.Println(userID)
		c.Set("userID", userID)
		c.Next()

	}
}
