package loggers

import (
	"fmt"
	"io"
	"log"
	"os"

	"ripflux/config"
	"sync"
	"time"
)

var (
	once       sync.Once
	mu         sync.Mutex
	logFile    *os.File
	baseLogger *log.Logger
	lastHeader string
)

func initLogger() {
	var err error
	logFile, err = os.OpenFile(config.COMBINED_LOG_FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	baseLogger = log.New(logFile, "", log.Ltime)
}

func Log(v ...any) {
	once.Do(initLogger)

	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	currentDate := now.Format("2006-01-02")

	if lastHeader != currentDate {
		lastHeader = currentDate
		header := fmt.Sprintf("\n========================= LOG SEGMENT: %s =========================\n", now.Format("Monday, 02 Jan 2006"))
		io.WriteString(logFile, header)

	}
	baseLogger.Println(v...)

}

func Logf(format string, v ...any) {
	Log(fmt.Sprintf(format, v...))
}

func mInitLogger(fileName string) {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)

	}

	baseLogger = log.New(logFile, "", log.Ltime)

}

func printInTerminal(v ...any) {
	fmt.Println(v...)
}

func TaskLog(fileName string, v ...any) {
	once.Do(func() {
		mInitLogger(fileName)
	})

	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	currentDate := now.Format("2006-01-02")

	if lastHeader != currentDate {
		lastHeader = currentDate
		header := fmt.Sprintf("\n========================= LOG SEGMENT: %s =========================\n", now.Format("Monday, 02 Jan 2006"))
		io.WriteString(logFile, header)

	}
	baseLogger.Println(v...)
}

func MTaskLog(fileName string, printOnTerminal bool, v ...any) error {
	once.Do(func() {
		mInitLogger(fileName)
	})

	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	currentDate := now.Format("2006-01-02")

	if lastHeader != currentDate {
		lastHeader = currentDate
		header := fmt.Sprintf("\n========================= LOG SEGMENT: %s =========================\n", now.Format("Monday, 02 Jan 2006"))
		io.WriteString(logFile, header)

	}
	baseLogger.Println(v...)
	if printOnTerminal == true {
		printInTerminal(v...)
	}

	return fmt.Errorf("ERROR: Check Logs")
}
