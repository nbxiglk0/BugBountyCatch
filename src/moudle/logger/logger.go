package logger

import (
	"fmt"
	"github.com/projectdiscovery/fileutil"
	"os"
	"path/filepath"
	"time"
)

var logFilePath string

func Logging(mes string) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		mes = time.Now().Format("\r\n2006-01-02 15:04:05") + ": " + mes
		_, err := file.WriteString(mes)
		if err != nil {
			fmt.Println("Log write fails" + err.Error())
		}
	}
}
func InitLogFile(homeDir string) {
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath = filepath.Join(homeDir, logFileName)
	flag := fileutil.FileExists(logFilePath)
	if flag == false {
		_, _ = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE, 0777)
	}
}
