package website

import (
	"github.com/gin-gonic/gin"
	sv "github.com/wilsonwu/golang-gin-project-layout/internal/services"
	"github.com/wilsonwu/golang-gin-project-layout/internal/services/website"
)

var Home = home{}

type home struct{}

func (c *home) HomePage(ctx *gin.Context) {
	s := sv.Context(ctx)
	users, _ := website.HomeService(ctx).UserList()
	s.View("website.home", gin.H{
		"Title": "首页 - ",
		"Users": users,
	})
}
