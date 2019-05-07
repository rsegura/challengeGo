package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"challenge/server"
)
var a server.App

func TestMain(m *testing.M) {
	a = server.App{}
	a.Init()
	code := m.Run()
	os.Exit(code)
}

func TestPokemonsEndpoints(t *testing.T) {

	req, _ := http.NewRequest("GET", "/pokemons", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an non empty array. Got %s", body)
	}
}
func TestAllEndpoint(t *testing.T){
	req, _ := http.NewRequest("GET", "/all", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetOnePokemon(t *testing.T) {
	req, _ := http.NewRequest("GET", "/pokemon/charmander", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an non empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

