package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Logger 新規 logrus インスタンスの生成
var Logger = logrus.New()

// InitLogger Logger の初期設定
func InitLogger(mode bool) {
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)

	// 環境変数にDEBUG_MODEが設定されていたら、Debugレベルのログまで出す
	if mode {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func FuncCostTime(path string, fn string) func() {
	start := time.Now()
	return func() {
		Logger.WithField("path", path).WithField("func", fn).Debugln("costTime", fmt.Sprintf(`%v`, time.Since(start)))
	}
}
