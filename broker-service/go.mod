module broker

go 1.18

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.0
)

require google.golang.org/protobuf v1.28.1 // indirect

require github.com/rabbitmq/amqp091-go v1.3.4 // direct
