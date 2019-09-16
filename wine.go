package main

import "sync"

// Vinho struct de um vinho
type Vinho struct {
	Nome        string  `json:"nome" valid:"required"`
	Safra       string  `json:"safra,omitempty" valid:"optional"`
	TipoVinho   string  `json:"tipoVinho" valid:"required"`
	TipoUva     string  `json:"tipoUva,omitempty" valida:"optional"`
	Volume      float64 `json:"volumeLitros,omitempty" valid:"optional"`
	QtdGarrafas int     `json:"qtdGarrafas,omitempty" valid:"optional"`
	Pais        string  `json:"pais,omitempty" valid:"optional"`
	Descricao   string  `json:"descricao,omitempty" valid:"optional"`
	Vinhedo     string  `json:"vinhedo,omitempty" valid:"optional"`
}

// Adega struct
type Adega struct {
	sync.Mutex
	A map[string]Vinho
}

// TempVinhos guarda um array de vinhos, para inserção em massa
type TempVinhos struct {
	V []Vinho `json:"vinhos"`
}

// Repository definição das funções que controlam a adega
type Repository interface {
	InitAdega() (*Adega, error)
	AddVinho(v *Vinho, a *Adega) error
	OverrideVinho(v *Vinho, a *Adega) error
	DeleteVinho(vinho string, a *Adega) error
	UpdateQtdGarrafasVinho(vinho string, qtd int, a *Adega) error
	BuscarPorNome(nome string, a *Adega) (*Vinho, error)
	ListaVinhos(a *Adega) ([]Vinho, error)
}
