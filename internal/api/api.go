package api

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Api struct {
	e *gin.Engine
}

func New() *Api {
	return &Api{}
}

func (a *Api) Start(context.Context) {
	err := a.e.Run(":8080")
	if err != nil {
		slog.Error("failed to start api")
	}
}

func (a *Api) Init() error {
	a.initServices()
	a.initRoutes()
	return nil
}

func (a *Api) initServices() {

}

func (a *Api) initRoutes() {

}
