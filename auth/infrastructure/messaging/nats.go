package messaging

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func ConnectToNATS(url string) (*nats.Conn, error) {
	opts := []nats.Option{nats.Name("Your NATS Client Name")}
	opts = append(opts, nats.ReconnectWait(time.Second))
	opts = append(opts, nats.MaxReconnects(10))

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
		return nil, err
	}
	return nc, nil
}
