package utils

import (
	"io"
	"log"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// LoggingSettings ログファイルの出力
func LoggingSettings(logFile string) error {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

	// logrusを用いたエラーメッセージの標準出力
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetLevel(logrus.WarnLevel)

	return nil
}
