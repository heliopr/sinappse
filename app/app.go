package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var JWTSecret string
var OAuth *oauth2.Config
var Server *gin.Engine

func LoadConfig() error {
	err := godotenv.Load("config/.env")
	if err != nil {
		return err
	}

	JWTSecret = os.Getenv("JWT_SECRET")

	OAuth = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		RedirectURL:  "http://localhost:9000/auth/login/callback/",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return nil
}

func InitDatabase() error {
	if DB != nil {
		return fmt.Errorf("database is already created")
	}
	err := ConnectToDatabase(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		return err
	}

	buf, err := os.ReadFile("config/init.sql")
	if err == nil {
		_, err := DB.Exec(string(buf))
		return err
	} else if !os.IsNotExist(err) {
		return err
	}
	
	return nil
}

func InitHttpServer() error {
	if Server != nil {
		return fmt.Errorf("server is already created")
	}

	Server = gin.Default()
	return nil
}

func RunHttpServer() error {
	return Server.Run(":"+os.Getenv("PORT"))
}


/*func getEnvInt(key string, defaultv int) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultv
	}
	return v
}*/