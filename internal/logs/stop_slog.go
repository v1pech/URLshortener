package logs

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

func Stop(logger *slog.Logger, logPath string, start_date time.Time) {
	fileNameIndex := strings.LastIndex(logPath, "/") + 1
	startTime := start_date.Format("01.02.2006 15:04")
	currentTime := time.Now().Format("01.02.2006 15:04")
	newLogPath := logPath[:fileNameIndex] + " " + startTime + "-" + currentTime + ".log"
	_ = os.Rename(logPath, newLogPath)
}
