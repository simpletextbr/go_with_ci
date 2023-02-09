package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(100, 233)

	if total != 333 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 333)
	}
}
