package models

type Aluno struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    CPF   string `json:"cpf"`
    RG    string `json:"rg"`
}