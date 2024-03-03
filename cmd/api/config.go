package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type limiter struct {
	rps     float64
	burst   int
	enabled bool
}

type db struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type smtp struct {
	host     string
	port     int
	username string
	password string
	sender   string
}

type config struct {
	port    int
	env     string
	db      db
	limiter limiter
	smtp    smtp
}

func validateAndConvertEnvVariableToInt(envVar string) int {
	v := validateStringEnvVariable(envVar)

	numericalEnvVariable, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("%s could not be converted to int", envVar)
	}

	return numericalEnvVariable
}

func validateAndConvertEnvVariableToBool(envVar string) bool {
	v := validateStringEnvVariable(envVar)

	boolValue, err := strconv.ParseBool(v)
	if err != nil {
		log.Fatalf("%s could not be converted to bool", envVar)
	}

	return boolValue
}

func validateStringEnvVariable(envVar string) string {
	envVariable := os.Getenv(envVar)

	if envVariable == "" {
		log.Fatalf("%s not configured in .env file", envVar)
	}

	return envVariable
}

func loadConfigFromEnvVariables() config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db{
		dsn:          validateStringEnvVariable("DB_CONN_STRING"),
		maxOpenConns: validateAndConvertEnvVariableToInt("DB_MAX_OPEN_CONN"),
		maxIdleConns: validateAndConvertEnvVariableToInt("DB_MAX_IDLE_CONN"),
		maxIdleTime:  validateStringEnvVariable("DB_MAX_IDLE_TIME"),
	}

	limiter := limiter{
		rps:     float64(validateAndConvertEnvVariableToInt("REQ_LIMITER_RPS")),
		burst:   validateAndConvertEnvVariableToInt("REQ_LIMITER_BURST"),
		enabled: validateAndConvertEnvVariableToBool("REQ_LIMITER_EMABLED"),
	}

	smtp := smtp{
		host:     validateStringEnvVariable("MAILER_HOST"),
		port:     validateAndConvertEnvVariableToInt("MAILER_PORT"),
		username: validateStringEnvVariable("MAILER_USERNAME"),
		password: validateStringEnvVariable("MAILER_PASSWORD"),
		sender:   validateStringEnvVariable("MAILER_SENDER"),
	}

	cfg := config{
		port:    validateAndConvertEnvVariableToInt("PORT"),
		env:     validateStringEnvVariable("ENV"),
		db:      db,
		limiter: limiter,
		smtp:    smtp,
	}

	return cfg
}
