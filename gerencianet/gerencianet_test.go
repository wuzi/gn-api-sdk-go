package gerencianet

import (
	"testing"
)

func TestNewGerencianet(t *testing.T) {
	params := map[string]interface{}{
		"client_id": "cidTest",
		"client_secret": "csTest",
		"sandbox": true,
		"timeout": 10,
	}
	gn := NewGerencianet(params)
	if gn == nil {
		t.Error("Error on constructor")
	}
}