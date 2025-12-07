package main

import (
	"log"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

const logFile = "morseDecoder.log"

func main() {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	log.SetOutput(file)

	srv := server.NewServer(logger)
	logger.Printf("Start server: %s\n", time.Now().UTC().String())

	err = srv.Server.ListenAndServe()
	if err != nil {
		log.Fatal("server start error", err)
	}

}
