package router

import (
	"be-golang-chapter-30/template-project-golang/database"
	"be-golang-chapter-30/template-project-golang/handler"
	"be-golang-chapter-30/template-project-golang/repository"
	"be-golang-chapter-30/template-project-golang/service"
	"be-golang-chapter-30/template-project-golang/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRouter() (*chi.Mux, *zap.Logger, error) {
	r := chi.NewRouter()

	logger := util.InitLog()

	config, err := util.ReadConfiguration()
	if err != nil {
		return nil, logger, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return nil, logger, err
	}

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	Handle := handler.NewAllHandler(service, logger, config)

	r.HandleFunc("/", Handle.SampelHandler.Create)

	return r, logger, nil
}
