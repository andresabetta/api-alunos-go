package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Saudacao(c *gin.Context) {
    nome := c.Param("nome")
    c.JSON(http.StatusOK, gin.H{
        "API diz": "E ai " + nome + ", Tudo beleza?",
    })
}