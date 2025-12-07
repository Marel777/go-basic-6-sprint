package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

const (
	serverAddr         = ":8080"          // Порт сервера
	serverReadTimeout  = 5 * time.Second  // Таймаут для чтения
	serverWriteTimeout = 10 * time.Second // Таймаут для записи
	serverIdleTimeout  = 15 * time.Second // Таймаут ожидания следующего запроса
)

// Структура с полями для логгера и http сервера
type MorseDecoder struct {
	Logger *log.Logger
	Server *http.Server
}

// Функция создает mux и возвращает экземпляр сервера
func NewServer(logger *log.Logger) MorseDecoder {
	var md MorseDecoder
	mux := http.NewServeMux()
	registerHandlers(mux)                 // Зарегистрировали обработчики
	md.Server = createServer(mux, logger) // Создали http
	md.Logger = logger                    // Логгер
	return md
}

// Функция регистрации Handlers
func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc(`/`, handlers.MainHandler)
	mux.HandleFunc(`/upload`, handlers.UploadHandler)
}

// Создание http сервера
func createServer(mux *http.ServeMux, logger *log.Logger) *http.Server {
	var server http.Server
	server.Addr = serverAddr
	server.Handler = mux
	server.ErrorLog = logger
	server.ReadTimeout = serverReadTimeout
	server.WriteTimeout = serverWriteTimeout
	server.IdleTimeout = serverIdleTimeout
	return &server
}
