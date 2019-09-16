package main

import (
	"errors"
	"log"
)

// InitAdega Inicializa a adega, vazia
func InitAdega() (*Adega, error) {
	var adega Adega

	adega.A = make(map[string]Vinho)

	return &adega, nil
}

// BuscarPorNome retorna o vinho buscando pelo nome
func BuscarPorNome(nome string, a *Adega) (*Vinho, error) {

	a.Lock()
	vinho, exists := a.A[nome]
	a.Unlock()

	if !exists {
		//Vinho não localizado
		return nil, errors.New("Vinho não localizado")
	}

	// Vinho localizado
	return &vinho, nil
}

// AddVinho adiciona vinho à adega
func AddVinho(v *Vinho, a *Adega) error {
	log.Println("Iniciando AddVinho")

	log.Println("Validando se o vinho existe")

	_, exists := BuscarPorNome(v.Nome, a)

	if exists == nil {
		log.Println("Vinho já existente")
		return errors.New("Vinho já existente")
	}

	log.Println("Adicionando vinho")
	a.Lock()
	a.A[v.Nome] = (*v)
	a.Unlock()

	log.Println("Finalizando AddVinho com sucesso")
	return nil
}

// OverrideVinho retira o vinho com o mesmo nome (se existir) e adiciona o novo
func OverrideVinho(v *Vinho, a *Adega) error {
	log.Println("Iniciando OverrideVinho")

	a.Lock()

	log.Println("Apagando antigo")
	delete(a.A, v.Nome)

	log.Println("Adicionando novo vinho")
	a.A[v.Nome] = (*v)

	a.Unlock()

	log.Println("Finalizando OverrideVinho")
	return nil
}

// DeleteVinho apaga vinho v da adega
func DeleteVinho(vinho string, a *Adega) error {
	log.Println("Iniciando DeleteVinho")

	_, exists := BuscarPorNome(vinho, a)

	if exists != nil {
		// Vinho não existente
		return errors.New("Vinho não existente")
	}

	a.Lock()

	delete(a.A, vinho)

	a.Unlock()

	log.Println("Vinho deletado")
	return nil
}

// UpdateQtdGarrafasVinho adiciona ou retira qtd do vinho
func UpdateQtdGarrafasVinho(vinho string, qtd int, a *Adega) error {
	log.Println("Iniciando UpdateQtdGarrafasVinho")

	v, exists := BuscarPorNome(vinho, a)

	if exists != nil {
		// Vinho não existente
		return errors.New("Vinho não existente")
	}
	a.Lock()

	qtdAtual := v.QtdGarrafas

	a.Unlock()

	if qtdAtual+qtd < 0 {
		return errors.New("Quantidade negativa, valor inválido")
	}

	log.Println("Atualizando quantidade")
	a.Lock()

	v.QtdGarrafas = v.QtdGarrafas + qtd
	a.A[vinho] = (*v)

	a.Unlock()

	log.Println("Finalizando UpdateQtdGarrafasVinho")
	return nil

}

// ListaVinhos retorna array de vinhos
func ListaVinhos(a *Adega) ([]Vinho, error) {

	var vinhos []Vinho
	a.Lock()
	for _, value := range a.A {
		vinhos = append(vinhos, value)
	}
	a.Unlock()

	if len(vinhos) == 0 {
		//Sem vinhos
		return nil, errors.New("Nenhum vinho")
	}

	return vinhos, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
