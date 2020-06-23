package service

import (
	"github.com/pkg/errors"

	"github.com/ifukuazumi/scrumwise_sample/model"
	"github.com/ifukuazumi/scrumwise_sample/usecase/repository"
)

// Production is
type Production interface {
	GetTagID() (string, error)
	GetScrumwise() ([]model.SprintBacklogs, error)
}

type productionImpl struct {
	ProductionRepository repository.Production
}

func NewProduction(productionRepository repository.Production) Production {
	return &productionImpl{ProductionRepository: productionRepository}
}

func (p *productionImpl) GetTagID() (string, error) {
	return p.ProductionRepository.GetTagID()
}

func (p *productionImpl) GetScrumwise() ([]model.SprintBacklogs, error) {
	data, err := p.ProductionRepository.GetAll()
	if err != nil {
		return nil, errors.WithMessage(err, "p.ProductionRepository.Get()でエラーが起きました")
	}

	backlogs := p.salvageBacklogs(data.Backlogs)

	var sprintBacklogs []model.SprintBacklogs
	for _, sprint := range data.Sprints {
		sprintBacklogs = append(sprintBacklogs, model.NewSprintBacklogs(sprint, backlogs))
	}

	return sprintBacklogs, nil
}

// salvageBacklogs TODO ここを上手い感じに書き換えましょう
func (*productionImpl) salvageBacklogs(backlogs []model.Backlog) []model.Backlog {
	var result []model.Backlog
	for _, backlog := range backlogs {
		result = append(result, backlog)
		if backlog.ChildBacklogItems == nil {
			continue
		}

		for _, child := range backlog.ChildBacklogItems {
			result = append(result, child)
			if child.ChildBacklogItems == nil {
				continue
			}
			for _, child2 := range child.ChildBacklogItems {
				result = append(result, child2)
			}
		}
	}
	return result
}
