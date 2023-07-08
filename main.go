package main

import (
	"fmt"
	"gin-gorm-microservice/infrastructure/repository/config"
	"gin-gorm-microservice/infrastructure/rest/controllers/errors"
	"gin-gorm-microservice/infrastructure/rest/routes"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())
	var err error
	db, err := config.GormOpen()
	if err != nil {
		_ = fmt.Errorf("fatal error in database file: #{err}")
		panic(err)
	}
	router.Use(errors.Handler)
	routes.ApplicationV1Router(router, db)
	startServer(router)
}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)
	}
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description : %s", strings.ToLower(err.Error()))
		panic(err)
	}
}
