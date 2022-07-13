package log

import (
	logr "github.com/go-logr/zerologr"
    "github.com/rs/zerolog"
	"sync"
	"gopkg.in/natefinch/lumberjack.v2"
)

var instance *Logger
var once sync.Once

type Logger struct {
	logr.Logger
}

func GetInstance(name string) *Logger {
	once.Do(func(){
		createInstance(name)
	})

	return instance
}

func createInstance(name string) {
	// zerolog.New()でなぜか怒られる
	// writer := lumberjack.Logger{
	// var writer io.Writer
	// writer = &lumberjack.Logger{
	// 		Filename: "/var/log/gwse/app.log", // output file
	// 	MaxSize: 10, // MB
	// 	// RollingInterval: 1, // day
	// 	MaxBackups: 10, // max quantity
	// 	LocalTime: true, // default UTC
	// 	Compress: true, // archive flag
	// }
	// z := zerolog.New(writer)

	z := zerolog.New(&lumberjack.Logger {
		Filename: "/var/log/es/app.log", // output file
		MaxSize: 10, // MB
		// RollingInterval: 1, // day
		MaxBackups: 10, // max quantity
		LocalTime: true, // default UTC
		Compress: true, // archive flag
	}).With().Caller().Timestamp().Logger()
	logger := logr.New(&z)
	instance = &Logger{logger}
}