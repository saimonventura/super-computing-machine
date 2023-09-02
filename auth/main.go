package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var ctx = context.Background()
var rdb *redis.Client

// Define the structure for input data
type User struct {
	UUID     string `json: "uuid"`
	Email    string `json:email"`
	Password string `json: "password"`
}

func initializeDevUser(rbd *redis.Client) {
	// Define the development user
	user := &User{
		UUID:     "UUIDc8930a08-0500-4dad-933e-6f4f519fc470",
		Email:    "dev@saimonventura.com",
		Password: "$2a$10$0pTKxeXBzuT.EUULMXaNaeOUGlcmAdzlCjjbCE8zu1wC7SotjY17q",
	}

	// Convert the user to a JSON string for storage
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	// Set the user in Redis, using the email as the key
	err = rbd.Set(ctx, user.Email, userJSON, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set user: %v", err)
	}
	fmt.Println("Development user initialized")
}

func InitializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	// Initialize the development user
	initializeDevUser(rdb)
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")[7:]
		claims := &jwt.MapClaims{}

		secretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add the user to the context
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq User

	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch stored user data from Redis
	storedUserJSON, err := rdb.Get(ctx, loginReq.Email).Result()
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var storedUser User
	err = json.Unmarshal([]byte(storedUserJSON), &storedUser)
	if err != nil {
		fmt.Println("Error unmarshalling user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check the provided password against the stored one
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginReq.Password))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		fmt.Println("JWT_SECRET_KEY not set")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": loginReq.UUID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// Handle error
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

func main() {
	// Initialize Redis connection
	InitializeRedis()

	r := mux.NewRouter()
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")

	http.Handle("/", r)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
