package helpers

import (
	"os"
)

func Env() string {
	val := os.Getenv("WAVE_ENV")
	if val == "" {
		return "production"
	}
	return val
}

func Production() bool {
	return Env() == "production"
}

func Development() bool {
	return Env() == "development"
}
