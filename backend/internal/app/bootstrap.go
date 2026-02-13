package app

import (
	backendhttp "safeguild/backend/internal/app/http"

	"github.com/gin-gonic/gin"
)

func BuildHTTPRouter(engine *gin.Engine, deps backendhttp.Dependencies) {
	backendhttp.Register(engine, deps)
}
