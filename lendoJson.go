package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Função para transformar a string em data
func parseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Erro ao analisar a data:", err)
	}
	return t
}

func lendoPessoaJson() {
	// Abrir o arquivo JSON
	file, err := os.Open("JSON/Pessoa.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	//Fecha o arquivo após a execução
	defer file.Close()

	// Ler o conteúdo do arquivo
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Copia o conteúdo lido do JSON para a estrutura Pessoa
	var pessoa Pessoa
	err = json.Unmarshal(byteValue, &pessoa)
	if err != nil {
		log.Fatalf("Erro ao copiar conteúdo Pessoa: %v", err)
	}
	//Chama a função que transforma a string em data
	pessoa.DataNascimento = parseDate(pessoa.DataNascimentoStr)

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
	fmt.Println("--------------------------------------------------")
}

func lendoEnderecosJson() {
	// Abrir o arquivo JSON
	file, err := os.Open("JSON/Enderecos.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Ler o conteúdo do arquivo
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	// Copia cada conteúdo lido do JSON para a estrutura Endereco
	var enderecos []Endereco
	err = json.Unmarshal(byteValue, &enderecos)
	if err != nil {
		log.Fatalf("Erro ao copiar conteúdo Endereco: %v", err)
	}
	//Laço for para percorrer o vetor de endereços
	for _, endereco := range enderecos {
		fmt.Printf("EnderecoId = %d\n", endereco.EnderecoId)
		fmt.Printf("PessoaId = %d\n", endereco.PessoaId)
		fmt.Printf("Logradouro = %s\n", endereco.Logradouro)
		fmt.Printf("Cidade = %s\n", endereco.Cidade)
		fmt.Printf("Bairro = %s\n", endereco.Bairro)
		fmt.Printf("Estado = %s\n", endereco.Estado)
		fmt.Printf("Numero = %s\n", endereco.Numero)
		fmt.Printf("Cep = %s\n", endereco.Cep)
		fmt.Printf("Ativo = %t\n", endereco.Ativo)
		fmt.Printf("DataUltAlt = %s\n", endereco.DataUltAlt.Format("2006-01-02 15:04:05"))
		fmt.Printf("DataCadastro = %s\n", endereco.DataCadastro.Format("2006-01-02 15:04:05"))
		fmt.Println("--------------------------------------------------")
	}
}
