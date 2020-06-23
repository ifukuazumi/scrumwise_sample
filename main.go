package main

import (
	_ "github.com/ChimeraCoder/gojson"
	"github.com/ifukuazumi/scrumwise_sample/adapter"
	"github.com/ifukuazumi/scrumwise_sample/infra"
	"github.com/ifukuazumi/scrumwise_sample/log"
	"github.com/ifukuazumi/scrumwise_sample/model"
	"github.com/ifukuazumi/scrumwise_sample/usecase/service"
	"os"
)

func main() {

	// 実行環境変数の読み込み
	appConfig, err := model.NewAppConfig()
	if err != nil {
		log.Logger.Fatal(err)
	}

	// ロガーの初期設定
	log.InitLogger(appConfig.DebugMode)

	// envファイル環境変数の読み込み
	if err := infra.LoadEnvFile(appConfig.AppEnv); err != nil {
		log.Logger.Fatal(err)
	}

	credentialUserName := os.Getenv("SCRUMWISE_CREDENTIAL_USER_NAME")
	credentialPassword := os.Getenv("SCRUMWISE_CREDENTIAL_PASSWORD")
	projectID := os.Getenv("PROJECT_ID")
	tagName := os.Getenv("TAG_NAME")
	productionRepository := adapter.NewScrumwise(credentialUserName, credentialPassword, projectID, tagName)

	productionService := service.NewProduction(productionRepository)
	tagID, err := productionService.GetTagID()
	if err != nil {
		log.Logger.Fatal(err)
	}
	result, err := productionService.GetScrumwise()
	if err != nil {
		log.Logger.Fatal(err)
	}

	for _, sprintBacklogs := range result {
		log.Logger.Println(sprintBacklogs.Sprint.Name, ", ", sprintBacklogs.TagCount(tagID), "個")
	}

}
