package useage

import (
	"fmt"
	"os"
	"utils/pkg/config"
	"utils/pkg/logger"
)

func All() {
	Logger()
	Config()
}

func Logger() {
	pwd, _ := os.Getwd()
	path := pwd + "/logs"
	logger.Start(logger.Options{LogLv: "info", RemainDay: 30, Path: path})
	logger.Info("Logger.Info:%v", "Logger")
	logger.Debug("Logger.Debug:%v", "Logger")
	logger.Error("Logger.Error:%v", "Logger")
}

func Config() {
	config.Start()
	fmt.Println("Config:", config.Config.GetString("test"), config.Config.GetString("testhost.ip"))
}
