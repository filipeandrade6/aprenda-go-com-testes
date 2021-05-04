package main

import "testing"

func TestOla(t *testing.T) {
	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		// Necessário para dizermos ao conjunto de testes que este é um método auxiliar
		// Ao fazer isso, quando o teste falhar, o número da linha relatada estará em
		// nossa chamada de função, e não dentro do nosso auxiliar de teste.
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Olá, Chris"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo', quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"
		verificarMensagemCorreta(t, resultado, esperado)
	})
}
