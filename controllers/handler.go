package controllers

import (
	
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