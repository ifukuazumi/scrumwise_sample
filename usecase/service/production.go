package service

import (
	"fmt"
	"github.com/ifukuazumi/scrumwise_sample/log"
	"github.com/ifukuazumi/scrumwise_sample/model"
	"github.com/ifukuazumi/scrumwise_sample/usecase/repository"
)

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
	backlogData, err := p.ProductionRepository.GetBacklogs()
	if err != nil {
		fmt.Sprintln("p.ProductionRepository.Get()でエラーが起きました")
	}

	sprintData, err := p.ProductionRepository.GetSprints()
	if err != nil {
		fmt.Sprintln("p.ProductionRepository.Get()でエラーが起きました")
	}

	//_, err = p.ProductionRepository.GetTags()
	//if err != nil {
	//	fmt.Sprintln("p.ProductionRepository.Get()でエラーが起きました")
	//}

	_, err = p.collectTasksForEachSprint(backlogData, sprintData)
	if err != nil {
		return err
	}

	return nil
}

func (p *productionImpl) collectTasksForEachSprint(backlogData, sprintData *model.Result) ([]model.Task, error) {
	//var allTasks []model.Task
	for _, projectItemForSprint := range sprintData.Projects {
		for _, sprintItem := range projectItemForSprint.Sprints {
			log.Logger.Println("sprint: ",sprintItem.Name, sprintItem.ID)
			var allTasks []model.Task
			for _, projectItemForBacklog := range backlogData.Projects {
				for _, backlogItem := range projectItemForBacklog.Backlogs {
					if sprintItem.ID == backlogItem.SprintID {
						allTasks = append(allTasks, backlogItem.Tasks...)
						//log.Logger.Println("task: ", len(allTasks), "個, sprintName: ",sprintItem.Name, ", backlogName: ",backlogItem.Name)
					}
					if len(backlogItem.ChildBacklogItems) > 0 {
						for _, child := range backlogItem.ChildBacklogItems {
							if sprintItem.ID == child.SprintID {
								allTasks = append(allTasks, child.Tasks...)
								//log.Logger.Println("task: ", len(allTasks), "個, sprintName: ",sprintItem.Name, ", childBacklogName: ",child.Name)
							}

						}
					}
				}
				log.Logger.Println("allTask: ", len(allTasks), "個, splintID: ", sprintItem.ID, ", sprintName: ",sprintItem.Name)
				//log.Logger.Println(allTasks)
				log.Logger.Println("=========================")
			}
		}
	}
	//return allTasks, nil
	return nil, nil
}