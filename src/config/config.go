package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MarvelConfig struct {
	Host string

	PubKey string
	PvtKey string
}

type Config struct {
	Marvel MarvelConfig
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func New() *Config {

	API_URL, _     := os.LookupEnv("MARVEL_API_URL")
	API_PVT_KEY, _ := os.LookupEnv("MARVEL_API_PVT_KEY")
	API_PUB_KEY, _ := os.LookupEnv("MARVEL_API_PUB_KEY")

	return &Config{
		Marvel: MarvelConfig{
			Host:  API_URL,

			PubKey: API_PUB_KEY,
			PvtKey: API_PVT_KEY,
		},
	}
}