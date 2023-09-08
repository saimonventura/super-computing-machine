// auth/infrastructure/messaging/nats_messenger.go

package messaging

import (
	"github.com/nats-io/nats.go"
)

type NatsMessenger struct {
	conn *nats.Conn
}

func NewNatsMessenger(url string) (*NatsMessenger, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsMessenger{conn: conn}, nil
}

func (n *NatsMessenger) Publish(subject string, message []byte) error {
	return n.conn.Publish(subject, message)
}

func (n *NatsMessenger) Subscribe(subject string, callback func(msg []byte)) error {
	_, err := n.conn.Subscribe(subject, func(m *nats.Msg) {
		callback(m.Data)
	})
	return err
}
