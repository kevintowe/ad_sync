package load_config

import "os"

// Config struct holds all the configuration values
type Config struct {
	MessageBrokerUser     string
	MessageBrokerPassword string
	MessageBrokerUrl      string
}

// LoadConfig loads configuration values from environment variables and returns a Config object
func LoadConfig() Config {
	// Load environment variables and fall back to defaults if necessary
	messageBrokerUser := getEnv("MESSAGE_BROKER_USER", "guest")
	messageBrokerPassword := getEnv("MESSAGE_BROKER_PASSWORD", "guest")
	messageBrokerUrl := getEnv("MESSAGE_BROKER_URL", "amqp://guest:guest@localhost:5672/")

	// Create and return the Config struct
	return Config{
		MessageBrokerUser:     messageBrokerUser,
		MessageBrokerPassword: messageBrokerPassword,
		MessageBrokerUrl:      messageBrokerUrl,
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
