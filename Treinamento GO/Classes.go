package main

import (
	"fmt"
	"time"
)

type Pessoa struct {
	PessoaId       int       `json:"pessoaId"`
	Nome           string    `json:"nome"`
	Sobrenome      string    `json:"sobrenome"`
	Sexo           string    `json:"sexo"`
	CPF            string    `json:"cpf"`
	Telefone       string    `json:"telefone"`
	Email          string    `json:"email"`
	DataNascimento time.Time `json:"dataNascimento"`
	Ativo          bool      `json:"ativo"`
	DataUltAlt     time.Time `json:"dataUltAlt"`
	DataCadastro   time.Time `json:"dataCadastro"`
}

type Endereco struct {
	enderecoId   int       `json:"enderecoId"`
	pessoaId     int       `json:"pessoaId"`
	logradouro   string    `json:"logradouro"`
	cidade       string    `json:"cidade"`
	bairro       string    `json:"bairro"`
	estado       string    `json:"estado"`
	numero       int       `json:"numero"`
	cep          string    `json:"cep"`
	ativo        bool      `json:"ativo"`
	detaUltAlt   time.Time `json:"dataUltAlt"`
	dataCadastro time.Time `json:"dataCadastro"`
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

func main() {
	// Example usage
	pessoa := Pessoa{
		PessoaId:       1,
		Nome:           "João",
		Sobrenome:      "Silva",
		Sexo:           "M",
		CPF:            "12345678901",
		Telefone:       "1234567890",
		Email:          "joao.silva@example.com",
		DataNascimento: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Ativo:          true,
		DataUltAlt:     time.Now(),
		DataCadastro:   time.Now(),
	}
	endereco := Endereco{
		enderecoId:   1,
		pessoaId:     pessoa.PessoaId,
		logradouro:   "Rua das Flores",
		cidade:       "São Paulo",
		bairro:       "Centro",
		estado:       "SP",
		numero:       123,
		cep:          "12345678",
		ativo:        true,
		detaUltAlt:   time.Now(),
		dataCadastro: time.Now(),
	}

	fmt.Printf("PessoaID = %d\n", pessoa.PessoaId)
	fmt.Printf("Nome = %s\n", pessoa.Nome)
	fmt.Printf("Sobrenome = %s\n", pessoa.Sobrenome)
	fmt.Printf("Sexo = %s\n", pessoa.Sexo)
	fmt.Printf("CPF = %s\n", pessoa.CPF)
	fmt.Printf("Telefone = %s\n", pessoa.Telefone)
	fmt.Printf("Email = %s\n", pessoa.Email)
	fmt.Printf("DataNascimento = %s\n", pessoa.DataNascimento.Format("2006-01-02"))
	fmt.Printf("Ativo = %t\n", pessoa.Ativo)
	fmt.Printf("DataUltAlt = %s\n", pessoa.DataUltAlt.Format("2006-01-02 15:04:05"))
	fmt.Printf("DataCadastro = %s\n", pessoa.DataCadastro.Format("2006-01-02 15:04:05"))

	fmt.Printf("////////////////////////////////////////////////////\n")

	fmt.Printf("enderecoId = %d\n", endereco.enderecoId)
	fmt.Printf("pessoaId = %d\n", endereco.pessoaId)
	fmt.Printf("logradouro = %s\n", endereco.logradouro)
	fmt.Printf("cidade = %s\n", endereco.cidade)
	fmt.Printf("bairro = %s\n", endereco.bairro)
	fmt.Printf("estado = %s\n", endereco.estado)
	fmt.Printf("numero = %d\n", endereco.numero)
	fmt.Printf("cep = %s\n", endereco.cep)
	fmt.Printf("ativo = %t\n", endereco.ativo)
	fmt.Printf("dataUltAlt = %s\n", endereco.detaUltAlt.Format("2006-01-02 15:04:05"))
	fmt.Printf("dataCadastro = %s\n", endereco.dataCadastro.Format("2006-01-02 15:04:05"))
}
