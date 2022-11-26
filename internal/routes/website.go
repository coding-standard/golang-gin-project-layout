package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsonwu/golang-gin-project-layout/internal/apps/website"
)

func RegisterWebsiteRoute(engine *gin.Engine) {
	group := engine.Group("/")

	group.GET("/", website.Home.HomePage)
}
