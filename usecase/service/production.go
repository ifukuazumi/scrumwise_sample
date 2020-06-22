package service

import (
	"fmt"

	"github.com/ifukuazumi/scrumwise_sample/log"
	"github.com/ifukuazumi/scrumwise_sample/model"
	"github.com/ifukuazumi/scrumwise_sample/usecase/repository"
)

const channelBuffer  = 5

// Production is
type Production interface {
	GetScrumwise() error
}

type productionImpl struct {
	ProductionRepository      repository.Production
}

func NewProduction (productionRepository repository.Production) Production {
	return &productionImpl{ProductionRepository:productionRepository}
}


func (p *productionImpl) GetScrumwise() error {
	data, err := p.ProductionRepository.GetAll()
	if err != nil {
		fmt.Sprintln("p.ProductionRepository.Get()でエラーが起きました")
	}

	var backlogs []model.Backlog
	for _, backlog := range data.Backlogs {
		backlogs = append(backlogs, backlog)
		if backlog.ChildBacklogItems == nil {
			continue
		}

		for _, child := range backlog.ChildBacklogItems {
			backlogs = append(backlogs, child)
			if child.ChildBacklogItems == nil {
				continue
			}
			for _, child2 := range child.ChildBacklogItems {
				backlogs = append(backlogs, child2)
			}
		}
	}

	tagID, err := p.ProductionRepository.GetTagID()
	if err != nil {
		fmt.Sprintln("p.ProductionRepository.Get()でエラーが起きました")
	}

	var sprintBacklogs []model.SprintBacklogs
	for _, sprint := range data.Sprints {
		splintID := sprint.ID
		tagCount := 0
		var targetBacklogs []model.Backlog
		for _, backlog := range backlogs {
			if splintID == backlog.SprintID {
				targetBacklogs = append(targetBacklogs, backlog)
			}
		}
		for _, sprintBacklog := range targetBacklogs {
			for _, task := range sprintBacklog.Tasks {
				for _, tagNums := range task.TagIDs {
					if tagNums == tagID {
						tagCount++
					} 
				}
			}
		}
		
		log.Logger.Println(sprint.Name, ", ", tagCount,"個")
		
		sprintBacklogs = append(sprintBacklogs, model.SprintBacklogs{
			Sprint:   sprint,
			Backlogs: targetBacklogs,
		})
	}


	return nil
}