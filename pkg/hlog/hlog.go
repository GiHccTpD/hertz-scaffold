package hlog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
	"time"

	hertzzap "github.com/hertz-contrib/logger/zap"
)

func init() {
	// Customizable output directory.
	var logFilePath string
	dir := "./"
	logFilePath = path.Join(dir, "/logs/")
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	// For zap detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/zap and https://github.com/uber-go/zap
	logger := hertzzap.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
}
