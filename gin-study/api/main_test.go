package main

import (
	"gin-study/api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/andre", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebeido foi %d e o esperado era %d", response.Code, http.StatusOK)
	}
}
