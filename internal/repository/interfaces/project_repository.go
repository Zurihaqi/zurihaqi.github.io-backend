package interfaces

import "zurihaqi.github.io-backend/internal/model"

type ProjectRepository interface {
	Create(project *model.Project) (*model.Project, error)
	FindAll() ([]model.Project, error)
	FindByID(id uint) (*model.Project, error)
	Update(project *model.Project) (*model.Project, error)
	Delete(project *model.Project) error
}
