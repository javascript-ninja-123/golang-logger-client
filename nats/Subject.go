package nats




type NatsEventType string

const (
	EntityUploaded NatsEventType = "EntityUploaded"
	LOGGING NatsEventType = "Logging"
)

func ParseNatsEventType(name NatsEventType) string {
	return string(name)
}