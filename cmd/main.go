package main

import (
	"github.com/wilsonwu/golang-gin-project-layout/cmd/server"
	_ "github.com/wilsonwu/golang-gin-project-layout/pkg/config"
	_ "github.com/wilsonwu/golang-gin-project-layout/pkg/db"
	//_ "github.com/wilsonwu/golang-gin-project-layout/pkg/redis"
)

func main() {
	server.Run()
}
