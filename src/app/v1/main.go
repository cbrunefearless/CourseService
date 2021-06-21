package main

import (
	"fmt"
	"log"
	"net/http"
	"CourseService/app/v1/service"
	"CourseService/config"
	"github.com/gorilla/mux"
	logger "CourseService/logger"
	authentication "CourseService/authentication"
)


func handleRequests(config config.Config) {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Methods(http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut).Subrouter()
	myRouter.HandleFunc("/v1/courses", service.ReturnAllCourses)
	myRouter.HandleFunc("/v1/course/{id}", service.ReturnSingleCourse)
	myRouter.HandleFunc("/v1/course", service.CreateNewCourse).Methods("POST")
	myRouter.HandleFunc("/v1/course", service.UpdateCourse).Methods("PUT")
	myRouter.HandleFunc("/v1/course", service.DeleteCourse).Methods("DELETE")
	myRouter.HandleFunc("/v1/course/health", service.HealthCheck)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", config.Port), myRouter))
	myRouter.Use(authentication.Middleware)
}

func main() {

	var standardLogger = logger.NewLogger()

	standardLogger.Info("Starting program")

	config, err := config.LoadConfig("../../config")
	if err != nil {
        log.Fatal("Cannot load configuration. Error: ", err)
	}
	fmt.Println(fmt.Sprint("Rest API v1.0 - Mux Routers Starting on localhost:", config.Port))
	handleRequests(config)
}
