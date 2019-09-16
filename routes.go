package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetAdega retorna JSON com todos os vinhos
func GetAdega(c *gin.Context) {
	log.Println("Iniciando rota GetAdega")

	vinhos, err := ListaVinhos(adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, vinhos)
	return
}

// PutVinho adiciona/atualiza vinho
func PutVinho(c *gin.Context) {
	log.Println("Iniciando rota PutVinho")

	var V Vinho
	err := c.BindJSON(&V)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validação de campos obrigatórios
	_, err = govalidator.ValidateStruct(V)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Realizando inserção/override do vinho
	err = OverrideVinho(&V, adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Vinho inserido/atualizado")
	c.JSON(201, gin.H{"message": "Vinho inserido/atualizado com sucesso"})
}

// PostVinho adiciona vinho
func PostVinho(c *gin.Context) {
	log.Println("Iniciando rota PostVinho")

	var V Vinho
	err := c.BindJSON(&V)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validação de campos obrigatórios
	_, err = govalidator.ValidateStruct(V)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Realizando inserção do vinho
	err = AddVinho(&V, adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Vinho inserido")
	c.JSON(201, gin.H{"message": "Vinho inserido com sucesso"})
}

// UpdateQtd atualiza qtd
func UpdateQtd(c *gin.Context) {
	log.Println("Iniciando rota UpdateQtd")

	var Vinho string
	Vinho = c.Param("vinho")
	var QtdS string
	QtdS = c.Param("quantidade")

	Qtd, err := strconv.Atoi(QtdS)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Quantidade: ", Qtd)
	// Realizando atualização da quantidade
	err = UpdateQtdGarrafasVinho(Vinho, Qtd, adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Quantidade atualizada")
	c.JSON(200, gin.H{"message": "Quantidade atualizada"})
}

// GetVinho retorna vinho
func GetVinho(c *gin.Context) {
	log.Println("Iniciando rota GetVinho")

	var Vinho string
	Vinho = c.Param("vinho")

	// Realizando atualização da quantidade
	V, err := BuscarPorNome(Vinho, adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna vinho
	log.Println("Retornando vinho: ", V.Nome)
	c.JSON(200, V)
}

// PostVinhos adiciona um array de vinhos
func PostVinhos(c *gin.Context) {
	log.Println("Iniciando rota PostVinhos")

	var V TempVinhos
	err := c.BindJSON(&V)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Realizando validações")
	for _, vinho := range V.V {
		// Validação de campos obrigatórios
		_, err = govalidator.ValidateStruct(vinho)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Valida se já existe
		_, err = BuscarPorNome(vinho.Nome, adega)
		if err == nil {
			erro := fmt.Sprintf("Vinho já existente: %s", vinho.Nome)
			log.Println(erro)
			c.JSON(http.StatusBadRequest, gin.H{"error": erro})
			return
		}
	}

	for _, vinho := range V.V {
		// Realizando inserção do vinho
		err = AddVinho(&vinho, adega)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	log.Println("Vinhos inseridos")
	c.JSON(201, gin.H{"message": "Vinhos inseridos com sucesso"})
}

// DelVinho apaga vinho caso exista
func DelVinho(c *gin.Context) {
	log.Println("Iniciando rota DeleteVinho")

	var Vinho string
	Vinho = c.Param("vinho")

	// Realizando atualização da quantidade
	err := DeleteVinho(Vinho, adega)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna vinho
	log.Println("Vinho apagado")
	c.JSON(200, gin.H{"message": "Vinho apagado com sucesso."})
}
