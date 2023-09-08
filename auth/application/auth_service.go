package application

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"super-computing-machine/auth/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	RedisClient      *redis.Client
	SecretKey        string
	messagingService domain.MessagingService
}

func NewAuthService(client *redis.Client, secretKey string, messaging domain.MessagingService) *AuthService {
	return &AuthService{
		RedisClient:      client,
		SecretKey:        secretKey,
		messagingService: messaging,
	}
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	ctx := context.Background()

	// Fetch stored user data from Redis
	storedUserJSON, err := s.RedisClient.Get(ctx, email).Result()
	if err != nil {
		return "", errors.New("unauthorized")
	}

	var storedUser domain.User
	err = json.Unmarshal([]byte(storedUserJSON), &storedUser)
	if err != nil {
		return "", errors.New("error unmarshalling user")
	}

	// Check the provided password against the stored one
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New("unauthorized")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": storedUser.UUID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", errors.New("could not generate token")
	}

	msg := []byte("Login successful for user: " + email)
	s.messagingService.Publish("user.login.success", msg)

	return tokenString, nil
}
