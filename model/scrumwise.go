package model

// Response is
type Response struct {
	Result Result `json:"result"`
}

// Result is
type Result struct {
	Projects []Project `json:"projects"`
}

// Project is
type Project struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Backlogs []Backlog `json:"backlogItems"`
	Sprints  []Sprint  `json:"sprints"`
	Tags     []Tag     `json:"tags"`
}

// Backlog is
type Backlog struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	SprintID          string    `json:"sprintID"`
	ChildBacklogItems []Backlog `json:"childBacklogItems"`
	Tasks             []Task    `json:"tasks"`
}

// Task is
type Task struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	TagIDs []string `json:"tagIDs"`
}

// Sprint is
type Sprint struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SprintBacklogs is
type SprintBacklogs struct {
	Sprint   Sprint
	Backlogs []Backlog
}

// Tag is
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
