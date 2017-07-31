package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/freeality/pet-amigo-server/controllers"
	"github.com/freeality/pet-amigo-server/db"
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
	iniciarDb()

	router.Run(":" + port)
}

// iniciarDb cria as tabelas durante o desenvolvimento
// as tabelas serão recriadas.
func iniciarDb() {
	db.Migrate()
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
