// auth/domain/messaging/messenger.go

package messaging

type Messenger interface {
	Publish(subject string, message []byte) error
	Subscribe(subject string, callback func(msg []byte)) error
}
