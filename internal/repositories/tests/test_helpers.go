package tests

import "github.com/joho/godotenv"

func settingUpEnvs() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
}