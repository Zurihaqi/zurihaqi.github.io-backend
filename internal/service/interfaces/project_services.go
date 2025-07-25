package interfaces

import "zurihaqi.github.io-backend/internal/model"

type ProjectService interface {
	CreateProject(project *model.Project) (*model.Project, error)
	FindAllProjects() ([]model.Project, error)
	FindProjectByID(id uint) (*model.Project, error)
	UpdateProject(id uint, projectData *model.Project) (*model.Project, error)
	DeleteProject(id uint) error
}
