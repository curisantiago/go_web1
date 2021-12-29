package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID              int
	Nombre          string  `json:"nombre" binding:"required"`
	Apellido        string  `json:"apellido" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Edad            int     `json:"edad" binding:"required"`
	Altura          float64 `json:"altura" binding:"required"`
	Activo          bool    `json:"activo" binding:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
}

var users []Request

var tokenPedido = "1319"

func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// verificacion de token
		token := ctx.GetHeader("token")
		if token != tokenPedido {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, fmt.Sprintf("el campo %s es requerido", err.Error()))
			return
		} else {
			if len(users) == 0 {
				req.ID = 1
			} else {
				req.ID = users[len(users)-1].ID + 1
			}
			users = append(users, req)
			ctx.JSON(200, req)
			fmt.Println(users)
		}

	}
}

func main() {
	r := gin.Default()

	r.POST("/usuarios/new", Guardar())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
