package filelogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type FileLogger struct {
	logger *log.Logger
}

func New(fileName string, directory string, flag int) (*FileLogger, error) {
	fileLogger := &FileLogger{}

	err := os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		return fileLogger, fmt.Errorf("failed to make log file directory: %w", err)
	}

	file, err := os.OpenFile(path.Join(directory, fileName), flag|os.O_CREATE|os.O_WRONLY, 0644)
	fileLogger.logger = log.New(file, "", log.LstdFlags)

	return fileLogger, fmt.Errorf("failed to open log file: %w", err)
}

func (fileLogger FileLogger) Println(msg ...any) {
	fileLogger.logger.Println(msg...)
}

func (fileLogger FileLogger) Print(msg ...any) {
	fileLogger.logger.Print(msg...)
}

func (fileLogger FileLogger) Disable() {
	fileLogger.logger.SetOutput(io.Discard)
}
