package service

import (
	"be-golang-chapter-30/template-project-golang/repository"

	"go.uber.org/zap"
)

type AllService struct {
	SampelService SampelService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		SampelService: NewSampelService(repo, log),
	}
}
