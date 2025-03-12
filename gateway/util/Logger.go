package util

import (
	"log"
	"os"
)

func LogInfof(s string, args ...any) {
	log.SetPrefix("[INFO] ")
	log.Printf(s, args...)
	log.SetPrefix("")
}

func LogWarnf(s string, args ...any) {
	log.SetPrefix("[WARN] ")
	log.Printf(s, args...)
	log.SetPrefix("")
}

func LogErrorf(s string, args ...any) {
	log.SetPrefix("[ERROR] ")
	log.Printf(s, args...)
	log.SetPrefix("")
}

func LogFatalf(s string, args ...any) {
	log.SetPrefix("[FATAL] ")
	log.Fatalf(s, args...)
	log.SetPrefix("")
}

func LogDebugf(s string, args ...any) {
	if os.Getenv("DEBUG") != "" {
		log.SetPrefix("[DEBUG] ")
		log.Printf(s, args...)
		log.SetPrefix("")
	}
}
