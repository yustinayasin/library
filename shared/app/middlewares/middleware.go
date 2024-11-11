package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/proto"
	"shared/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId int `json:"id"`
	jwt.MapClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

type JWTConfig struct {
	Claims                  *JwtCustomClaims
	SigningKey              []byte
	ErrorHandlerWithContext func(error, *gin.Context) error
}

type GeneratorToken interface {
	GenerateToken(userId int) string
}

func (jwtConf *ConfigJWT) Init() JWTConfig {
	return JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c *gin.Context) error {
			utils.ErrorResponseWithoutMessages(c, http.StatusForbidden, e.Error())
			return nil
		},
	}
}

func (configJWT ConfigJWT) GenerateToken(userId int) string {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(configJWT.ExpiresDuration)).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(configJWT.SecretJWT))

	return token
}

func verifyToken(tokenString string, jwtConf ConfigJWT, c *gin.Context) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(jwtConf.SecretJWT), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

// Auth for private routes
func RequireAuth(next gin.HandlerFunc, jwtConf ConfigJWT, userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookies
		// tokenString, err := c.Cookie("auth_token")

		// Get the token from header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You need to login first")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Invalid Authorization header format")
			c.Abort()
			return
		}

		// if err != nil {
		// 	fmt.Println(err)
		// 	utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You need to login first")
		// 	c.Abort()
		// 	return
		// }

		// Verify token
		token, err := verifyToken(parts[1], jwtConf, c)
		if err != nil {
			fmt.Println(err)
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}

		// Check the expiry date
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to extract claims")
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Token expired")
			c.Abort()
			return
		}

		// Extract userId from claims
		subsClaim, ok := claims["userId"].(float64)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to parse user ID")
			c.Abort()
			return
		}

		subsInt := int(subsClaim)

		// Get user from gRPC service
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		userResponse, err := userClient.GetUser(ctx, &proto.UserIdRequest{Id: int32(subsInt)})
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to fetch user data")
			c.Abort()
			return
		}

		c.Set("user", userResponse.User)

		next(c)
	}
}

func RequireAuthAdmin(next gin.HandlerFunc, jwtConf ConfigJWT, userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You need to login first")
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// Verify token
		token, err := verifyToken(tokenString, jwtConf, c)
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}

		// Check the expiry date
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to extract claims")
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Token expired")
			c.Abort()
			return
		}

		// Extract userId from claims
		subsClaim, ok := claims["userId"].(float64)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to parse user ID")
			c.Abort()
			return
		}

		subsInt := int(subsClaim)

		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		userResponse, err := userClient.GetUser(ctx, &proto.UserIdRequest{Id: int32(subsInt)})
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to fetch user data")
			c.Abort()
			return
		}

		// Check if user has admin role
		if userResponse.User.RoleId != 1 {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You don't have the authorization")
			c.Abort()
			return
		}

		c.Set("user", userResponse.User)

		next(c)
	}
}

func CacheControl(maxAge int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
		c.Next()
	}
}
