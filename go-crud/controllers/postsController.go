package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud/initializers"
	"go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Retirar dados do corpo do requerimento
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Criar um post (insert)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Devolver

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Receber o envio
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Responder
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Obter ID do URL
	id := c.Param("id")
	// Obter o posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Responder
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Capturar id do url
	id := c.Param("id")

	// Capturar os dados do requerimento do body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Encontrar o post quando atualizar
	var post models.Post
	initializers.DB.First(&post, id)

	// Atualizar
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Responder
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Capturar o id do URL
	id := c.Param("id")

	// Deletar o registro
	initializers.DB.Delete(&models.Post{}, id)

	// Responder
	c.JSON(200, gin.H{
		"Mensagem": "Registro deletado",
	})
}
