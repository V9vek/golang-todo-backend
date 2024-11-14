package utils

import "github.com/joho/godotenv"

func LoadEnvFile() {
	err := godotenv.Load()
	PanicIfError(err)
}
