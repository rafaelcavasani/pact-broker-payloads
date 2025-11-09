package pact

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProviderPact(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "rpaas",
		Provider: "hub",
	}
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:5002",
		BrokerURL:       "http://localhost:9292",
		ProviderVersion: "1.0.0",
	})
}
