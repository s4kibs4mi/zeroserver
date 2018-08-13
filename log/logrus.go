package log

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/s4kibs4mi/zeroserver/log/hooks"
)

var defLogger *logrus.Logger

func Init() {
	defLogger = logrus.New()
	defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())
}

func Logger() *logrus.Logger {
	return defLogger
}
