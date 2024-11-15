package service

import (
	"be-golang-chapter-30/template-project-golang/model"
	"be-golang-chapter-30/template-project-golang/repository"

	"go.uber.org/zap"
)

type SampelService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewSampelService(repo repository.AllRepository, Log *zap.Logger) SampelService {
	return SampelService{
		Repo:   repo,
		Logger: Log,
	}
}

func (sampelService *SampelService) Create(sampel *model.Sampel) error {
	return sampelService.Create(sampel)
}

func (sampelService *SampelService) GetAll() ([]*model.Sampel, error) {
	return nil, nil
}

func (sampelService *SampelService) GetByID(id int) (*model.Sampel, error) {
	return nil, nil
}

func (sampelService *SampelService) Update(id int, book model.Sampel) error {
	return nil
}

func (sampelService *SampelService) Delete(id int) error {
	return nil
}
