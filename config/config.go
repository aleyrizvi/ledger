package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port uint
	DBRW string
	DBRO string
}

func New() *Config {
	port := os.Getenv("PORT")
	portUint, err := strconv.ParseUint(port, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing PORT: %v", err)
		return nil
	}

	c := &Config{
		Port: uint(portUint),
		DBRW: os.Getenv("DBRW"),
		DBRO: os.Getenv("DBRO"),
	}

	return c
}
