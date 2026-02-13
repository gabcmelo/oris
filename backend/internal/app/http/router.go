package httpapp

import (
	legacy "oris/backend/internal/http/router"

	"github.com/gin-gonic/gin"
)

type Handlers = legacy.Handlers
type Dependencies = legacy.Dependencies

func Register(engine *gin.Engine, deps Dependencies) {
	legacy.Register(engine, legacy.Dependencies(deps))
}
