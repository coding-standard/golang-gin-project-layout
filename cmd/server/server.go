package server

import (
	"github.com/gin-contrib/static"
	"log"

	"github.com/wilsonwu/golang-gin-project-layout/pkg/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/wilsonwu/golang-gin-project-layout/internal/routes"
	"github.com/wilsonwu/golang-gin-project-layout/pkg/utils"
)

func Run() {
	engine := gin.Default()

	engine.SetFuncMap(utils.GetTemplateFuncMap())
	engine.Use(static.Serve("/", static.LocalFile("assets", false)))
	engine.LoadHTMLGlob("views/**/**/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	routes.RegisterWebsiteRoute(engine)

	if err := engine.Run(":8081"); err != nil {
		log.Fatalf("server running error: %v", err)
	}
}
