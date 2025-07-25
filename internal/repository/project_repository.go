package repository

import (
	"gorm.io/gorm"
	"zurihaqi.github.io-backend/internal/model"
	repoInterface "zurihaqi.github.io-backend/internal/repository/interfaces"
)

type projectRepositoryImpl struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repoInterface.ProjectRepository {
	return &projectRepositoryImpl{DB: db}
}

func (r *projectRepositoryImpl) Create(project *model.Project) (*model.Project, error) {
	err := r.DB.Create(project).Error
	return project, err
}

func (r *projectRepositoryImpl) FindAll() ([]model.Project, error) {
	var projects []model.Project
	err := r.DB.Find(&projects).Error
	return projects, err
}

func (r *projectRepositoryImpl) FindByID(id uint) (*model.Project, error) {
	var project model.Project
	err := r.DB.Where("id = ?", id).First(&project).Error
	return &project, err
}

func (r *projectRepositoryImpl) Update(project *model.Project) (*model.Project, error) {
	err := r.DB.Save(project).Error
	return project, err
}

func (r *projectRepositoryImpl) Delete(project *model.Project) error {
	return r.DB.Delete(project).Error
}
