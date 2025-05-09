package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/andresabetta/api-alunos/internal/handlers"
    "github.com/andresabetta/api-alunos/internal/models"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/saudacao/:nome", handlers.Saudacao)
    r.POST("/alunos", func(c *gin.Context) {
        c.Set("isTest", true)
        handlers.CreateAluno(c)
    })
    return r
}

func TestSaudacaoStatusCode(t *testing.T) {
    router := setupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/saudacao/gui", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"API diz":"E ai gui, Tudo beleza?"}`, w.Body.String())
}

func TestCreateAluno(t *testing.T) {
    router := setupRouter()
    // Remova esta linha: router.POST("/alunos", handlers.CreateAluno)

    aluno := models.Aluno{Nome: "João", CPF: "12345678901", RG: "123456789"}
    jsonValue, _ := json.Marshal(aluno)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/alunos", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    assert.Contains(t, w.Body.String(), `"message":"Aluno criado com sucesso"`)
}