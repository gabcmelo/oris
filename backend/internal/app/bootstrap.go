package app

import (
	backendhttp "oris/backend/internal/app/http"

	"github.com/gin-gonic/gin"
)

func BuildHTTPRouter(engine *gin.Engine, deps backendhttp.Dependencies) {
	backendhttp.Register(engine, deps)
}
