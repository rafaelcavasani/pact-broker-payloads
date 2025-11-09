package pact

import (
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestConsumerPact(t *testing.T) {
	// Configuração do Pact
	pact := &dsl.Pact{
		Consumer: "rpaas",
		Provider: "hub",
	}

	// Definição da interação
	pact.
		AddInteraction().
		Given("Provider is ready").
		UponReceiving("A request to calculate processing").
		WithRequest(dsl.Request{
			Method: "POST",
			Path:   dsl.String("/processing/v1/calculate"),
			Body: map[string]interface{}{
				"cliente": map[string]string{
					"codigo_pessoa": "12345678901",
				},
				"pagamento": map[string]string{
					"valor":          "12345678901",
					"data_pagamento": "2024-06-15",
				},
				"contrato": map[string]string{
					"numero_contrato": "9876543210",
				},
			},
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body: map[string]interface{}{
				"cliente": map[string]string{
					"codigo_pessoa": "12345678901",
				},
			},
		})

	// Verificação do contrato
	err := pact.Verify(func() error {
		// Simula a chamada ao Provider
		resp, err := http.Post("http://localhost:5002/processing/v1/calculate", "application/json", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	})

	if err != nil {
		t.Fatalf("Error verifying pact: %v", err)
	}
}
