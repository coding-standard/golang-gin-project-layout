package website

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wilsonwu/golang-gin-project-layout/internal/entities"
	"github.com/wilsonwu/golang-gin-project-layout/internal/models"
	"github.com/wilsonwu/golang-gin-project-layout/internal/services"
)

func HomeService(ctx *gin.Context) *sHome {
	return &sHome{ctx: services.Context(ctx)}
}

type sHome struct {
	ctx *services.BaseContext
}

func (s *sHome) UserList() ([]*entities.User, error) {
	var users []*models.User
	models.UserInit().M.Find(&users)
	var eusers []*entities.User
	for _, u := range users {
		var eu = entities.User{}
		err := copier.Copy(&eu, &u)
		if err != nil {
			return nil, err
		}
		eusers = append(eusers, &eu)
	}

	return eusers, nil
}
