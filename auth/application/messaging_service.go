package application

type MessagingService interface {
	Publish(event interface{}) error
}
