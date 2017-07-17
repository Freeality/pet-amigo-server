package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UsuarioController contém recursos para usuários
type UsuarioController struct { }

// Login retorna nil se o usuário estiver correto
func (c UsuarioController) Login(g *gin.Context) {
	g.HTML(http.StatusOK, "index.tmpl.html", gin.H {"Nome": "Usuário"}, )
}