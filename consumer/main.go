package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Estruturas fornecidas
type Cliente struct {
	CodigoPessoa string `json:"codigo_pessoa"`
}

type Pagamento struct {
	Valor         float64 `json:"valor"`
	DataDagamento string  `json:"data_pagamento"`
}

type PagamentoOutput struct {
	Valor             float64 `json:"valor"`
	DataDagamento     string  `json:"data_pagamento"`
	ValorDesconto     float32 `json:"valor_desconto"`
	DataProcessamento string  `json:"data_processamento"`
}

type Contrato struct {
	NumeroContrato string `json:"numero_contrato"`
}

type RequestPayload struct {
	Cliente   Cliente   `json:"cliente"`
	Pagamento Pagamento `json:"pagamento"`
	Contrato  Contrato  `json:"contrato"`
}

type ResponsePayload struct {
	Cliente   Cliente         `json:"cliente"`
	Pagamento PagamentoOutput `json:"pagamento"`
	Contrato  Contrato        `json:"contrato"`
}

// Função para chamar o provider
func callProvider(url string, payload RequestPayload) (*ResponsePayload, error) {
	// Serializa o payload para JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar o payload: %w", err)
	}

	// Cria a requisição HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	// Define os cabeçalhos
	req.Header.Set("Content-Type", "application/json")

	// Cria um cliente HTTP
	client := &http.Client{Timeout: 10 * time.Second}

	// Envia a requisição
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro: status code %d", resp.StatusCode)
	}

	// Decodifica a resposta JSON
	var responsePayload ResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	return &responsePayload, nil
}

func main() {
	// URL do provider
	providerURL := "http://localhost:5002/processing/v1/calculate"

	// Cria o payload de exemplo
	payload := RequestPayload{
		Cliente: Cliente{
			CodigoPessoa: "12345",
		},
		Pagamento: Pagamento{
			Valor:         100.50,
			DataDagamento: "2023-10-01",
		},
		Contrato: Contrato{
			NumeroContrato: "CONTRATO123",
		},
	}

	// Chama o provider
	response, err := callProvider(providerURL, payload)
	if err != nil {
		fmt.Printf("Erro ao chamar o provider: %v\n", err)
		return
	}

	// Exibe a resposta
	fmt.Printf("Resposta do provider: %+v\n", response)
}
