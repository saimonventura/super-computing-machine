// auth/domain/messaging.go

package domain

type MessagingService interface {
	Publish(subject string, message []byte) error
}
