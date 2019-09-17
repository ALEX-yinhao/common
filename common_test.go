package common_test

import (
	config "github.com/ALEX-yinhao/common"
	"testing"
)

func TestName(t *testing.T) {

	config.LoadConfig("./config_pre.toml")

	config.InitLog(config.GetMustStringValue("log", "LogDir"), config.GetMustStringValue("log", "LogPrefix"), config.GetMustStringValue("log", "LogSuffix"), config.GetMustIntValue("log", "LogSize"))

	config.Log.Info("hhh")

	config.Log.Info(config.Config)

}
