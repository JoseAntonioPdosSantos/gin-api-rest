package main

import (
	"encoding/json"
	"gin-api-rest/controller"
	"gin-api-rest/database"
	"gin-api-rest/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestCheckGreetingStatusCode(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controller.Greeting)
	req, _ := http.NewRequest("GET", "/gui", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

	mockResponse := `{"API say:":"hey there!gui, it's all right:"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}

func TestFindAllHandler(t *testing.T) {
	database.Connection()
	r := SetupTestRoutes()
	r.GET("/students", controller.GetAll)
	req, _ := http.NewRequest("GET", "/students", nil)
	request := httptest.NewRecorder()
	r.ServeHTTP(request, req)
	assert.Equal(t, http.StatusOK, request.Code)
}

func TestFindOneByCpf(t *testing.T) {
	database.Connection()
	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controller.GetOneByCpf)
	req, _ := http.NewRequest("GET", "/students/cpf/333.333.333-33", nil)
	request := httptest.NewRecorder()
	r.ServeHTTP(request, req)
	assert.Equal(t, http.StatusOK, request.Code)
}

func TestFindByIdHandle(t *testing.T) {
	database.Connection()
	r := SetupTestRoutes()
	r.GET("/students/:id", controller.GetOneByID)
	pathSearch := "/students/" + strconv.Itoa(5)
	req, _ := http.NewRequest("GET", pathSearch, nil)
	request := httptest.NewRecorder()
	r.ServeHTTP(request, req)
	var student model.Student
	json.Unmarshal(request.Body.Bytes(), &student)
	assert.Equal(t, "João Feitosa", student.Name)
	assert.Equal(t, "333.333.333-33", student.CPF)
}
