package main

import (
	"github.com/julienschmidt/httprouter"
	"mongodb-api/config"
	"mongodb-api/controllers"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(config.GetSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/user/:id", uc.UpdateUser)
	http.ListenAndServe(":8080", r)
}
