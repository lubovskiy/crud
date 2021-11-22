package main

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/lubovskiy/crud/infrastructure/database"
	"github.com/lubovskiy/crud/internal/config"
	"github.com/spf13/viper"


	"github.com/lubovskiy/crud/pkg/api"
)

func main() {
	ctx := context.Background()

	connAddress := config.GetConfigConnection()
	conn, err := database.NewConn(connAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}


	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}