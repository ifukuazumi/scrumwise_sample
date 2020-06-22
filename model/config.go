package model

import "github.com/kelseyhightower/envconfig"

// AppConfig プログラム実行時に指定する環境変数の構造体を用意
type AppConfig struct {
	//APP_ENVに指定がない場合、local
	AppEnv string `envconfig:"APP_ENV" default:"local"`
}

// NewAppConfig プログラム実行時に指定する環境変数を読み込み
func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
