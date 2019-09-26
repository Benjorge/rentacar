package veiculo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//owner(dono) do handler(manipulador)
type Controller struct {
	storage MySQLStorage
}

//construtor do nosso controller
func NewVeiculo(stg MySQLStorage) *Controller {

	return &Controller{
		storage: stg,
	}
}

//endpoint que busca os veiculos
//o que está dentro do 1 parenteses é o dono da função
func (ctrl *Controller) Get(c *gin.Context) {

	veiculos, err := ctrl.storage.GetVeiculos()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, veiculos)

}

//endpoint que cria novos veiculos
func (ctrl *Controller) Create(c *gin.Context) {
	var v Veiculo
	//tranforma a request em um objeto do tipo veiculo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	//salva os dados no banco
	err := ctrl.storage.CreateVeiculo(v.Nome, v.Marca, v.Ano, v.Modelo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// atualiza veiculos
func (ctrl *Controller) Update(c *gin.Context) {
	var v Veiculo

	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	//salva os dados no banco
	err := ctrl.storage.UpdateVeiculo(v.ID, &v)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)

}

//Apaga um veiculo
func (ctrl *Controller) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	//declara a variavel e ao mesmo tempo verifica se é diferente de nil
	if err := ctrl.storage.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)

}
