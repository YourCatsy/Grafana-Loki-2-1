package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"webapp/pkg/config"
)

func main() {
	port := config.GetEnv("PORT", "3000")

	// Настройка логирования
	logFilePath := "/var/log/webapp.log"
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Создание Echo
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))
	e.Use(middleware.Recover())

	// Роутинг
	e.Static("/", "public")
	e.GET("/", func(c echo.Context) error {
		log.Println("GET / request received")
		return c.File("public/views/webapp.html")
	})

	// Запуск сервера
	e.Logger.Fatal(e.Start(":" + port))
}
