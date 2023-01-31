package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Logging(mes string) {
	logFilePath := logFile()
	file, err := os.OpenFile(logFilePath, os.O_APPEND, 0644)
	if err == nil {
		mes = time.Now().Format("\r\n2006-01-02 15:04:05") + ": " + mes
		_, err := file.WriteString(mes)
		if err != nil {
			fmt.Print("Log write fails")
		}
	}
}
func logFile() string {
	logFileName := time.Now().Format("2006-01-02") + ".log"
	path, err := os.Getwd()
	if err != nil {
	}
	logFilePath := filepath.Join(path, logFileName)
	_, err = os.Stat(logFilePath)
	if err == nil {
	}
	if os.IsNotExist(err) {
		_, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE, 0644)
	}
	return logFilePath
}
