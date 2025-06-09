package main

import (
	"fmt"
	"os"
	"termbook/config"
	_ "termbook/docs"
	"termbook/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title TermBook API
// @version 1.0
// @description TermBook 的后端接口文档
// @host localhost:8080
// @BasePath /

func main() {
	viper.SetConfigFile("boot.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg config.Configure
	viper.Unmarshal(&cfg)
	os.Mkdir(config.DataPath, 0777)
	srv := gin.Default()
	srv.Use(cors.Default())
	routes.RegisterRoutes(srv,
		[]routes.Route{
			&routes.DocumentRoute{},
			&routes.RunnerRoute{},
			&routes.SystemRoute{},
		})
	srv.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
}
