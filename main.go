package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simplify-go/app"
	"simplify-go/controller"
	"simplify-go/exception"
	"simplify-go/helper"
	"simplify-go/middleware"
	"simplify-go/repository"
	"simplify-go/service"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	envErr := godotenv.Load(".env")
	helper.PanicIfError(envErr)

	webPort := os.Getenv("PORT")

	log.Printf("starting service on port %s\n", webPort)

	db := app.NewDB()
	validate := validator.New()

	repository := repository.NewExampleRepositoryImpl()

	service := service.NewExampleServiceImpl(repository, db, validate)

	controller := controller.NewExampleControllerImpl(service)

	router := httprouter.New()

	router.POST("/api/example", controller.Create)
	router.PUT("/api/example", controller.Update)
	router.DELETE("/api/example/:exampleId", controller.Delete)
	router.GET("/api/example/:exampleId", controller.FindById)
	router.GET("/api/example", controller.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
