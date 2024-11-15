package handler

import (
	"be-golang-chapter-30/template-project-golang/service"
	"be-golang-chapter-30/template-project-golang/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	SampelHandler SampelHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		SampelHandler: NewSampelService(service, log, config),
	}

}
