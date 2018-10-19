package main

import (
	"fmt"
	"time"

	"github.com/ExPreman/go-svg2pdf/middleware"
	"github.com/ExPreman/go-svg2pdf/pdf/delivery/http"
	"github.com/ExPreman/go-svg2pdf/pdf/usecase"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.InitMiddleware().CORS)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := usecase.NewPDFUsecase(timeoutContext)
	http.NewPDFHttpHandler(e, au)

	e.Start(viper.GetString("server.address"))
}
