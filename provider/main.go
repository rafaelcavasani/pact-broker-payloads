package main

import (
	"encoding/json"
	"net/http"
)

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

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var requestPayload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	responsePayload := ResponsePayload{
		Cliente: requestPayload.Cliente,
		Pagamento: PagamentoOutput{
			Valor:             requestPayload.Pagamento.Valor,
			DataDagamento:     requestPayload.Pagamento.DataDagamento,
			ValorDesconto:     5.0,
			DataProcessamento: "2024-06-20",
		},
		Contrato: requestPayload.Contrato,
	}

	// Retorna a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responsePayload)
}

func main() {
	http.HandleFunc("/processing/v1/calculate", calculateHandler)
	http.ListenAndServe(":5002", nil)
}
