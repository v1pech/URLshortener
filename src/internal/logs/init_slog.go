package logs

import (
	"io"
	"log"
	"log/slog"
	"os"
)

func Init(cfgLogPath, cfgLogLevel string) (*slog.Logger, *os.File) {

	logFile := openLogsFile(cfgLogPath)
	logOutput := io.MultiWriter(os.Stdout, logFile)

	var LogLevel slog.Level

	switch cfgLogLevel {
	case "debug":
		LogLevel = slog.LevelDebug
	case "info":
		LogLevel = slog.LevelInfo
	case "warn":
		LogLevel = slog.LevelWarn
	case "error":
		LogLevel = slog.LevelError
	}

	logger := slog.New(
		slog.NewJSONHandler(logOutput, &slog.HandlerOptions{Level: LogLevel}),
	)
	logger.Info("logger has been initialized")

	if logFile == nil {
		logger.Warn("using only stdout for logs")
	}

	return logger, logFile
}

func openLogsFile(Path string) *os.File {
	if Path == "false" {
		return nil
	}
	logFile, err := os.OpenFile(Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("can't open logs file")
			return nil
		}
	}
	return logFile
}
