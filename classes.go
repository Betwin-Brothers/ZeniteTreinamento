package main

import (
	"time"
)

type Pessoa struct {
	PessoaId          int       `json:"pessoaId"`
	Nome              string    `json:"nome"`
	Sobrenome         string    `json:"sobrenome"`
	Sexo              string    `json:"sexo"`
	CPF               string    `json:"cpf"`
	Telefone          string    `json:"telefone"`
	Email             string    `json:"email"`
	DataNascimentoStr string    `json:"dataNascimento"`
	DataNascimento    time.Time `json:"-"`
	Ativo             bool      `json:"ativo"`
	DataUltAlt        time.Time `json:"dataUltAlt"`
	DataCadastro      time.Time `json:"dataCadastro"`
}

type Endereco struct {
	EnderecoId   int       `json:"enderecoId"`
	PessoaId     int       `json:"pessoaId"`
	Logradouro   string    `json:"logradouro"`
	Cidade       string    `json:"cidade"`
	Bairro       string    `json:"bairro"`
	Estado       string    `json:"estado"`
	Numero       int       `json:"numero"`
	Cep          string    `json:"cep"`
	Ativo        bool      `json:"ativo"`
	DataUltAlt   time.Time `json:"dataUltAlt"`
	DataCadastro time.Time `json:"dataCadastro"`
}

type Profissional struct {
	ProfissionalId   int       `json:"profissionalId"`
	PessoaId         int       `json:"pessoaId"`
	Profissao        string    `json:"profissao"`
	EspecialidadeId  int       `json:"especialidadeId"`
	NumeroConselho   int       `json:"numeroConselho"`
	Email            string    `json:"email"`
	InicioExpediente time.Time `json:"inicioExpediente"`
	FinalExpediente  time.Time `json:"finalExpediente"`
	AlmocoInicio     time.Time `json:"almocoInicio"`
	AlmocoFinal      time.Time `json:"almocoFinal"`
	Ativo            bool      `json:"ativo"`
	DataUltAlt       time.Time `json:"dataUltAlt"`
	DataCadastro     time.Time `json:"dataCadastro"`
}
