package service

import (
	"zurihaqi.github.io-backend/internal/model"
	repoInterface "zurihaqi.github.io-backend/internal/repository/interfaces"
	"zurihaqi.github.io-backend/internal/service/interfaces"
)

type projectService struct {
	ProjectRepo repoInterface.ProjectRepository
}

func NewProjectService(repo repoInterface.ProjectRepository) interfaces.ProjectService {
	return &projectService{ProjectRepo: repo}
}

func (s *projectService) CreateProject(project *model.Project) (*model.Project, error) {
	return s.ProjectRepo.Create(project)
}

func (s *projectService) FindAllProjects() ([]model.Project, error) {
	return s.ProjectRepo.FindAll()
}

func (s *projectService) FindProjectByID(id uint) (*model.Project, error) {
	return s.ProjectRepo.FindByID(id)
}

func (s *projectService) UpdateProject(id uint, data *model.Project) (*model.Project, error) {
	existing, err := s.ProjectRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Title = data.Title
	existing.Description = data.Description
	existing.Repo = data.Repo
	existing.Live = data.Live

	return s.ProjectRepo.Update(existing)
}

func (s *projectService) DeleteProject(id uint) error {
	return s.ProjectRepo.Delete(&model.Project{ID: id})
}
