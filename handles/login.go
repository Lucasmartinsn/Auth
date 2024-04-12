package handles

import (
	tipo "authentication-system/data_type"
	models "authentication-system/model"
	services "authentication-system/services"

	"github.com/gin-gonic/gin"
)

func Verificalogin(c *gin.Context) {
	var login tipo.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"Mensage": "Decode do json",
			"Error":   err.Error(),
		})
		return
	}

	response, err := models.Verificalogin(string(login.Username), string(login.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"Mensage": "Usuario não existe.",
			"Erro": err.Error(),
		})
		return
	}
	token, err := services.NewJWTService().GenerateToken(response.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"token": "Token não gerado! retorno null de dados",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
