package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/andresabetta/api-alunos/internal/database" // Adiciona o pacote database
    "github.com/andresabetta/api-alunos/internal/models"  // Adiciona o pacote models
)

func Saudacao(c *gin.Context) {
    nome := c.Param("nome")
    c.JSON(http.StatusOK, gin.H{
        "API diz": "E ai " + nome + ", Tudo beleza?",
    })
}

func CreateAluno(c *gin.Context) {
    var aluno models.Aluno
    if err := c.ShouldBindJSON(&aluno); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Nos testes, usamos CreateAlunoMock
    if c.GetBool("isTest") {
        if err := database.CreateAlunoMock(aluno); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    } else {
        db, err := database.Connect()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar ao banco de dados"})
            return
        }
        defer db.Close()

        if err := database.CreateAluno(db, aluno); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Aluno criado com sucesso"})
}