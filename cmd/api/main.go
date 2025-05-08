package main

import (
    "github.com/gin-gonic/gin"
    "github.com/andresabetta/api-alunos/internal/database"
    "github.com/andresabetta/api-alunos/internal/handlers"
    "log"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }
    defer db.Close()

    r := gin.Default()
    r.GET("/saudacao/:nome", handlers.Saudacao)
    if err := r.Run(":8081"); err != nil {
        log.Fatal("Erro ao iniciar o servidor:", err)
    }
}