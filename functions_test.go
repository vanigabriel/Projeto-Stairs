package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdega(t *testing.T) {
	adega, _ = InitAdega()

	v := Vinho{
		Nome:        "Campo Largo",
		QtdGarrafas: 1,
	}

	// Insere vinho corretamente
	err := AddVinho(&v, adega)
	assert.Nil(t, err)

	// Insere novamente o mesmo vinho
	err = AddVinho(&v, adega)
	assert.NotNil(t, err)

	// Apaga vinho não existente
	err = DeleteVinho("Teste", adega)
	assert.NotNil(t, err)

	// Atualiza qtd corretamente
	err = UpdateQtdGarrafasVinho("Campo Largo", 1, adega)
	assert.Nil(t, err)

	// Atualiza qtd negativa
	err = UpdateQtdGarrafasVinho("Campo Largo", -100, adega)
	assert.NotNil(t, err)

	// Busca nome existente
	_, err = BuscarPorNome("Campo Largo", adega)
	assert.Nil(t, err)

	// Busca nome não existente
	_, err = BuscarPorNome("Teste", adega)
	assert.NotNil(t, err)

	// Apaga vinho corretamente
	err = DeleteVinho("Campo Largo", adega)
	assert.Nil(t, err)

}
