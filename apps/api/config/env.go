package config
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port             string
	GoogleClientID   string
	GoogleSecret     string
	GoogleRedirect   string
	JwtSecret        string
	FE_URL           string
}

var Envs Env

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env not found, using system environment")
	}

	Envs = Env{
		Port:           os.Getenv("PORT"),
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSecret:   os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirect: os.Getenv("GOOGLE_REDIRECT_URL"),
		JwtSecret:      os.Getenv("JWT_SECRET_KEY"),
		FE_URL:         os.Getenv("FE_URL"),
	}
}