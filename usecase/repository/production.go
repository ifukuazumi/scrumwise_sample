package repository

import "github.com/ifukuazumi/scrumwise_sample/model"

// Production scrumwise用のインターフェース
type Production interface {
	GetBacklogs() (*model.Result, error)
	GetSprints() (*model.Result, error)
	GetTags() (*model.Result, error)
}