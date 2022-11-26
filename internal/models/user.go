package models

import (
	"github.com/wilsonwu/golang-gin-project-layout/pkg/db"
	"gorm.io/gorm"
)

type User struct {
	Model
	Name string `json:"name"`
}

type UserModel struct {
	M     *gorm.DB
	Table string
}

func UserInit() *UserModel {
	return &UserModel{M: db.DB.Model(&User{}), Table: "user"}
}
