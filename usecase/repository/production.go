package repository

import "github.com/ifukuazumi/scrumwise_sample/model"

// Production scrumwise用のインターフェース
type Production interface {
	GetTagID() (string, error)
	GetAll() (*model.Project, error)
}