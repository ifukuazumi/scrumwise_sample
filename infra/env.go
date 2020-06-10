package infra

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadEnvFile 環境変数の設定
func LoadEnvFile(envName string) error {
	if err := godotenv.Load(fmt.Sprintf(".env/%s.env", envName)); err != nil {
		return err
	}
	return nil
}
