package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.

*/

type Usuario struct {
	Id              int
	Nombre          string
	Apellido        string
	Email           string
	Edad            int
	Altura          float64
	Activo          bool
	FechaDeCreacion string
}

var usuarios []Usuario

func readJSON() {
	readfile, erro := os.ReadFile("./usuario.json")
	if erro != nil {
		fmt.Println("error leyendo el json")
	}
	if err := json.Unmarshal(readfile, &usuarios); err != nil {
		fmt.Println(err)
	}
}

func getAll(contexto *gin.Context) {
	readJSON()
	//res, _ := json.Marshal(usuarios)
	//fmt.Println(res)

	query_values := contexto.Request.URL.Query()
	if len(query_values) == 0 {
		contexto.JSON(200, usuarios)
	} else {
		filtrados := filter(query_values, usuarios)
		contexto.JSON(200, filtrados)

	}
	//contexto.JSON(200, string(res))

	//contexto.JSON(200, usuarios)

}

func filter(query_values map[string][]string, usuarios []Usuario) []Usuario {
	res := []Usuario{}
	for i, values := range query_values {
		switch i {
		case "id":
			for j := range usuarios {
				fmt.Println(usuarios[j].Id)
				v, _ := strconv.Atoi(values[0])
				if usuarios[j].Id == v {
					res = append(res, usuarios[j])
				}
			}
		case "nombre":
			for _, user := range usuarios {
				if user.Nombre == values[0] {
					res = append(res, user)
				}
			}

		case "apellido":
			for _, user := range usuarios {
				if user.Apellido == values[0] {
					res = append(res, user)
				}
			}
		}

	}

	//fmt.Println(i)
	//fmt.Println(values)
	//fmt.Println(values[0])

	return res
}

func newHandler(ctx *gin.Context) {
	readJSON()
	found := false
	user := ctx.Param("id")
	id, _ := strconv.Atoi(user)
	fmt.Println(id)
	for _, usuario := range usuarios {
		fmt.Println(usuario.Id)
		if usuario.Id == id {
			found = true
			fmt.Println(found)
			ctx.JSON(200, usuario)
		}
	}
	if !found {
		ctx.JSON(404, "not found")
	}

}

func main() {
	r := gin.Default()

	//------------------ Ejercicio 2 ------------------
	r.GET("saludar/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")

		ctx.JSON(200, gin.H{
			"message": "Hola " + nombre,
		})
	})

	//------------------ Ejercicio 3 ------------------
	r.GET("/usuarios", func(c *gin.Context) {
		getAll(c)
	})

	// ------------------------------------------------------

	r.GET("usuarios/:id", func(ctx *gin.Context) {
		newHandler(ctx)
	})

	//------------------------------------------------------

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//new_r := gin.Default()
	//	r.GET("usuarios/:id", func(ctx *gin.Context) {
	//		newHandler(ctx)
	//	})

	//new_r.Run()
}
