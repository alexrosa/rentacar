package veiculo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//owner do handler
type Controller struct {
	storage Storage
}

//constructor do nosso controller
func NewVeiculo(stg Storage) *Controller {
	return &Controller{
		storage: stg,
	}
}

//endpoint que busca os veiculos
func (ctrl *Controller) Get(c *gin.Context) {
	veiculos, err := ctrl.storage.GetVeiculos()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, veiculos)
}

//endpoint que cria novos veiculos
func (ctrl *Controller) Create(c *gin.Context) {
	var v Veiculo
	//transforma a request em um objeto do tipo Veiculo
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

//atualiza veiculos
func (ctrl *Controller) Update(c *gin.Context) {
	var v Veiculo
	//transforma a request em um objeto do tipo Veiculo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	//salva os dados no banco
	err := ctrl.storage.UpdateVeiculo(v.ID, &v)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

//apaga um veiculo
func (ctrl *Controller) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	//declara a variavel e ao mesmo tempo verifica se é diferente de nil
	if err := ctrl.storage.DeleteVeiculo(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
