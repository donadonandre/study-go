package main

import (
	"gin-study/api/controllers"
	"gin-study/api/database"
	"gin-study/api/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno teste",
		CPF:  "12345678900",
		RG:   "123456789"}

	database.DB.Create(aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/andre", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

	mockDaResposta := `{"api-response":"E aí andre, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock() // Quando acabar toda a função, roda esse comando

	r := SetupRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678900", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
