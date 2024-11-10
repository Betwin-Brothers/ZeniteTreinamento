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
		EnderecoId:   1,
		PessoaId:     pessoa.PessoaId,
		Logradouro:   "Rua das Flores",
		Cidade:       "São Paulo",
		Bairro:       "Centro",
		Estado:       "SP",
		Numero:       123,
		Cep:          "12345678",
		Ativo:        true,
		DataUltAlt:   time.Now(),
		DataCadastro: time.Now(),
	}

	profissional := Profissional{
		ProfissionalId:   1,
		PessoaId:         1,
		Profissao:        "Médico",
		EspecialidadeId:  101,
		NumeroConselho:   12345,
		Email:            "medico@example.com",
		InicioExpediente: time.Now(),
		FinalExpediente:  time.Now().Add(8 * time.Hour),
		AlmocoInicio:     time.Now().Add(4 * time.Hour),
		AlmocoFinal:      time.Now().Add(5 * time.Hour),
		Ativo:            true,
		DataUltAlt:       time.Now(),
		DataCadastro:     time.Now(),
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

	fmt.Println("EnderecoId:", endereco.EnderecoId)
	fmt.Println("PessoaId:", endereco.PessoaId)
	fmt.Println("Logradouro:", endereco.Logradouro)
	fmt.Println("Cidade:", endereco.Cidade)
	fmt.Println("Bairro:", endereco.Bairro)
	fmt.Println("Estado:", endereco.Estado)
	fmt.Println("Numero:", endereco.Numero)
	fmt.Println("Cep:", endereco.Cep)
	fmt.Println("Ativo:", endereco.Ativo)
	fmt.Println("DataUltAlt:", endereco.DataUltAlt)
	fmt.Println("DataCadastro:", endereco.DataCadastro)

	fmt.Printf("////////////////////////////////////////////////////\n")

	fmt.Println("ProfissionalId:", profissional.ProfissionalId)
	fmt.Println("PessoaId:", profissional.PessoaId)
	fmt.Println("Profissao:", profissional.Profissao)
	fmt.Println("EspecialidadeId:", profissional.EspecialidadeId)
	fmt.Println("NumeroConselho:", profissional.NumeroConselho)
	fmt.Println("Email:", profissional.Email)
	fmt.Println("InicioExpediente:", profissional.InicioExpediente)
	fmt.Println("FinalExpediente:", profissional.FinalExpediente)
	fmt.Println("AlmocoInicio:", profissional.AlmocoInicio)
	fmt.Println("AlmocoFinal:", profissional.AlmocoFinal)
	fmt.Println("Ativo:", profissional.Ativo)
	fmt.Println("DataUltAlt:", profissional.DataUltAlt)
	fmt.Println("DataCadastro:", profissional.DataCadastro)
}
