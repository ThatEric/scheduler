package main

import (
	"fmt"
	"scheduler/processors"
	"scheduler/types"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	Lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// Logger
	var logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = &Lumberjack.Logger{
		Filename: viper.GetString("RunningLogFile"),
		MaxSize:  5,
	}

	// Scheduler
	scheduler := types.NewScheduler()
	processor := processors.NewProcessor(scheduler.ProcessorCh)

	httpProcessor := processors.HTTPProcessor{Name: "http", RunningLog: logger}
	//httpsProcessor := processors.HTTPProcessor{Name: "https", IsSSL: true}
	processor.AddProcessors(httpProcessor)
	processor.StartProcessing()

	scheduler.StartScheduling()
	scheduler.WatchForNewSchedules(true)

	fmt.Scanln()
}
