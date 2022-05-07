package controller

import (
	"gin-api-rest/database"
	"gin-api-rest/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetOneByID(c *gin.Context) {
	var student model.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func GetOneByCpf(c *gin.Context) {
	cpf := c.Param("cpf")
	student := model.Student{CPF: cpf}
	database.DB.Where(&student).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say:": "hey there!" + name + ", it's all right:",
	})
}

func Create(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	var student model.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "student successful deleted"})
}

func Update(c *gin.Context) {
	var student model.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func ShowIndexPage(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func PageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
