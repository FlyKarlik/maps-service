package config

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config of service
type Config struct {
	ServiceName         string
	ServerHost          string
	NatsHost            string
	DbDsn               string
	DbDnsSP             string
	DbPassFile          string
	DbSpPassFile        string
	JwtSecret           string
	AesSecret           string
	SentryDSN           string
	JaegerHost          string
	KafkaBrokers        string
	KafkaAutoOffsetRest string
	KafkaRequestTopic   string
	KafkaResponseTopic  string
	ParentServerHost    string
	RedisAddr           string
	RedisStyledMapsDB   int
}

// NewConfig generate new config
func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("NewConfig godotenv load failed")
	}

	redisStyledMapsDBString := os.Getenv("REDIS_STYLED_MAP_DB")
	redisStyledMapsDB, err := strconv.Atoi(redisStyledMapsDBString)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	return &Config{
		ServiceName:         os.Getenv("SERVICE_NAME"),
		ServerHost:          os.Getenv("SERVER_HOST"),
		NatsHost:            os.Getenv("NATS_HOST"),
		DbDsn:               os.Getenv("DB_DSN"),
		DbDnsSP:             os.Getenv("DB_DNS_SP"),
		DbPassFile:          os.Getenv("DB_PASSWORD_FILE"),
		DbSpPassFile:        os.Getenv("DB_SP_PASSWORD_FILE"),
		JwtSecret:           os.Getenv("JWT_SECRET"),
		AesSecret:           os.Getenv("AES_SECRET"),
		SentryDSN:           os.Getenv("SENTRY_DSN"),
		JaegerHost:          os.Getenv("JAEGER_HOST"),
		KafkaBrokers:        os.Getenv("KAFKA_BROKERS"),
		KafkaAutoOffsetRest: os.Getenv("KAFKA_AUTO_OFFSET_RESET"),
		KafkaRequestTopic:   os.Getenv("KAFKA_REQUEST_TOPIC"),
		KafkaResponseTopic:  os.Getenv("KAFKA_RESPONSE_TOPIC"),
		ParentServerHost:    os.Getenv("PARENT_SERVER_HOST"),
		RedisAddr:           os.Getenv("REDIS_ADDR"),
		RedisStyledMapsDB:   redisStyledMapsDB,
	}, nil
}

func (c *Config) GetDbDnsSp() (string, error) {
	var b bytes.Buffer

	pass, err := os.ReadFile(c.DbSpPassFile)
	if err != nil {
		return "", err
	}

	b.WriteString(c.DbDnsSP)
	b.WriteString(" ")
	b.WriteString("password=")
	b.WriteString(string(pass))

	return b.String(), nil
}

func (c *Config) GetDbDns() (string, error) {
	var b bytes.Buffer

	pass, err := os.ReadFile(c.DbPassFile)
	if err != nil {
		return "", err
	}

	b.WriteString(c.DbDsn)
	b.WriteString(" ")
	b.WriteString("password=")
	b.WriteString(string(pass))

	return b.String(), nil
}
