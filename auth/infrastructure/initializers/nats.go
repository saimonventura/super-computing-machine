// auth/infrastructure/initializers/nats.go

package initializers

import (
	"log"
	"super-computing-machine/auth/infrastructure/config"
	"super-computing-machine/auth/infrastructure/messaging"
)

func InitNats() *messaging.NatsMessenger {
	natsURL := config.GetEnvWithDefault("NATS_URL", "nats://localhost:4222")
	natsMessenger, err := messaging.NewNatsMessenger(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	return natsMessenger
}
