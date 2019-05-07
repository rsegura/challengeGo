package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type HandlerInterface interface{
	GetPing(c *gin.Context)
	GetPokemons(c *gin.Context)
	GetAllValues(c *gin.Context)
	GetAllPokemons(c *gin.Context)
	GetPokemon(c *gin.Context)
	GetAllClash(c *gin.Context)
	GetClash(c *gin.Context)
}

type Handler struct {
	
}

func NewHandler() (HandlerInterface, error) {
	return &Handler{
	}, nil
}


type HandlerWithoutFrameworkInterface interface{
	GetPing(w http.ResponseWriter, req *http.Request)
	GetAllPokemons(w http.ResponseWriter, req *http.Request)
	GetAllValues(w http.ResponseWriter, req *http.Request)
	GetPokemon(w http.ResponseWriter, req *http.Request)
	/*GetPokemons()
	GetAllValues()
	GetAllPokemons()
	GetPokemon()
	GetAllClash()
	GetClash()*/
}

type HandlerWithoutFramework struct {
	
}

func NewHandlerWithoutFramework() (HandlerWithoutFrameworkInterface, error) {
	return &HandlerWithoutFramework{
	}, nil
}