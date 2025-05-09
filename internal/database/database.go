package database

import (
    "database/sql"
    "errors"
    _ "github.com/lib/pq"
    "github.com/andresabetta/api-alunos/internal/models"
)

func Connect() (*sql.DB, error) {
    connStr := "host=localhost port=5432 user=aluno password=senha123 dbname=alunodb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

func CreateAlunoMock(aluno models.Aluno) error {
    // Simula a inserção no banco sem conectar ao PostgreSQL
    if aluno.Nome == "" || aluno.CPF == "" || aluno.RG == "" {
        return errors.New("dados inválidos")
    }
    return nil
}

func CreateAluno(db *sql.DB, aluno models.Aluno) error {
    query := `INSERT INTO alunos (nome, cpf, rg) VALUES ($1, $2, $3)`
    _, err := db.Exec(query, aluno.Nome, aluno.CPF, aluno.RG)
    if err != nil {
        return err
    }
    return nil
}