package filelogger

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type FileLogger struct {
	logger *log.Logger
}

func New(fileName string, directory string, flag int) (FileLogger, error) {
	var fileLogger FileLogger

	os.MkdirAll(directory, os.ModePerm)
	file, err := os.OpenFile(path.Join(directory, fileName), flag|os.O_CREATE|os.O_WRONLY, 0644)
	fileLogger.logger = log.New(file, "", log.LstdFlags)

	return fileLogger, err
}

func (fileLogger FileLogger) Println(msg ...interface{}) {
	fileLogger.logger.Println(msg...)
}

func (fileLogger FileLogger) Print(msg ...interface{}) {
	fileLogger.logger.Print(msg...)
}

func (fileLogger FileLogger) Disable() {
	fileLogger.logger.SetOutput(ioutil.Discard)
}
