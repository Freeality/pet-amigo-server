package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// UsuarioController contém recursos para usuários
type UsuarioController struct { }

var instance *UsuarioController
var once sync.Once

func GetUsuarioController() *UsuarioController {
	once.Do(func() {
		instance = &UsuarioController{}
	})
	return instance
}

// Login retorna nil se o usuário estiver correto
func (c *UsuarioController) NewLogin(g *gin.Context) {
	g.HTML(http.StatusOK, "login.tmpl.html", gin.H {"Nome": "Usuário"}, )
}

func (c *UsuarioController) ValidateLogin(g *gin.Context) {
	username := g.PostForm("username")
	password := g.PostForm("password")

	if len(password) > 0 {
		g.HTML(http.StatusOK, "index.tmpl.html", gin.H {"Username" : username})
		return
	}
	g.String(http.StatusUnauthorized, "A senha não pode ser vazia")
}
