package main

import (
	"tutorial/handler"
	"tutorial/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Rutas de Admin
	admin:=r.Group("admin")
	admin.POST("/login", handler.LoginHandler)
	admin.Use(middleware.ValidateToken())
	// Ruta admin/evento
	evento:=admin.Group("/eventos")
	evento.Use(middleware.Authorization([]int{1}))
	// Metodos
	evento.GET("/",handler.GetAll)
	evento.GET("/estado/:estado",middleware.Authorization([]int{3}), handler.GetEventoEstado)
	evento.POST("/",middleware.Authorization([]int{3}), handler.AddEvento)
	evento.DELETE("/:titulo",middleware.Authorization([]int{3}),handler.DeleteEvento)
	admin.PUT("/:titulo",middleware.Authorization([]int{3}),handler.UpdateEvento)


	//Rutas de Usuario
	user := r.Group("/user")
	user.GET("/eventos/:titulo",handler.GetEvento)
	user.GET("/eventos/fecha/:fecha",handler.GetEventof)
	user.GET("/eventos",handler.GetPublicados)
	user.GET("/eventos/asistir",handler.Getasistir)


	r.Run("localhost:8080")
}
