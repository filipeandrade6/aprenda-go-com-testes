package carteira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarteira(t *testing.T) {
	tests := []struct {
		name      string
		carteira  Carteira
		depositar Bitcoin
		retirar   Bitcoin
		expected  Bitcoin
		err       error
	}{
		{"depositar", Carteira{}, Bitcoin(10), Bitcoin(0), Bitcoin(10), nil},
		{"retirar", Carteira{saldo: Bitcoin(20)}, Bitcoin(0), Bitcoin(10), Bitcoin(10), nil},
		{"retirar com saldo insuficiente", Carteira{Bitcoin(20)}, Bitcoin(0), Bitcoin(100), Bitcoin(20), ErroSaldoInsuficiente},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.carteira.Depositar(tc.depositar)
			err := tc.carteira.Retirar(tc.retirar)

			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, tc.err, err.Error())
			}
			assert.Equal(t, tc.expected, tc.carteira.Saldo())
		})
	}
}
