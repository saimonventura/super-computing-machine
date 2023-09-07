package dbseed

import (
	"context"
	"encoding/json"
	"log"
	"super-computing-machine/auth/domain"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Seeder struct {
	rdb *redis.Client
	log domain.Logger
}

func NewSeeder(rdb *redis.Client, log domain.Logger) *Seeder {
	return &Seeder{rdb: rdb, log: log}
}

func (s *Seeder) InitializeDevUser() {
	// Define the development user
	user := &domain.User{
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
	err = s.rdb.Set(ctx, user.Email, userJSON, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set user: %v", err)
	}
	s.log.Debug("Development user initialized")
}
