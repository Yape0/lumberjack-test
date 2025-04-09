package main

import (
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/natefinch/lumberjack"
	//"golang.org/x/exp/slog"
	"os"
)

func initLogger(serviceName string) *slog.Logger {
	// Ensure the /logs directory exists
	_ = os.MkdirAll("logs", os.ModePerm)

	logPath := "logs/" + serviceName + ".log"

	//function lumberjack persisnya
	writer := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // 1 MB
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false, //yang asli true
	}

	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

// function that generate logs
func generateLargeLogs(logger *slog.Logger) {
	for i := 0; i < 2363; i++ {
		logger.Info("simulated log entry",
			slog.String("timestamp", time.Now().Format(time.RFC3339)),
			slog.String("level", "INFO"),
			slog.String("inquiry_id", strconv.Itoa(rand.Intn(1000000))),
			slog.String("msg", "This is a very long message to fill up the log file quickly. Padding: "+randomString(200)),
		)
	}
}

// function that generate random string
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	logger := initLogger("web_automation")
	generateLargeLogs(logger)
}
