package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/freeality/pet-amigo-server/controllers"
)

func main() {
	iniciarServidor()
}

func iniciarServidor() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	defineRotas(router)

	router.Run(":" + port)
}

func defineRotas(r *gin.Engine) {
	usuarioController := controllers.GetUsuarioController()

	r.GET("/", home)
	r.GET("/login", usuarioController.NewLogin)
	r.POST("/validateLogin", usuarioController.ValidateLogin)
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H {} )
}
