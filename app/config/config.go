package config

import (
	"github.com/joho/godotenv"
	"go-course/app/exception"
	"os"
)

type Config interface {
	Get(key string) string
}

type config struct {

}

func (config *config) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &config{}
}
