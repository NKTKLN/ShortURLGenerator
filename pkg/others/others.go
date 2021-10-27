package others

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserReg struct {
	Email string `json:"email"`
	Password string `json:"pass"`
}

type Url struct {
	Id string
	Url string
	Visits int
}

type ApiArchive struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Visits int `json:"visits"`
}

type ApiUrl struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

type Api struct {
	Type string `json:"type"`
}

func ErrorChecking(err error) {
    if err != nil {
        log.Print(err)
    }
}
// Text encoding 
func Encoding(text string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(text), 5)
	ErrorChecking(err)
    return string(bytes)
}

func GetInfromationFromEnv(file, name string) string {
    ErrorChecking(godotenv.Load(file))
	return os.Getenv(name)
}
