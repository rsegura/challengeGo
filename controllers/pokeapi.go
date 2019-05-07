package controllers

import(
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type ErrorProfile struct{
	Message string
	Error error
}

func(u *Handler) GetAllPokemons(c *gin.Context){
	var data interface{}
	jsonErr := getJson("https://pokeapi.co/api/v2/pokemon?limit=10000", &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c.JSON(500, gin.H{"message": "Error to retrieve pokemons", "error": jsonErr})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"pokemons": data})
	return
}

func(u *HandlerWithoutFramework) GetAllPokemons(w http.ResponseWriter, req *http.Request){
	var data pokemonRequest
	jsonErr := getJson("https://pokeapi.co/api/v2/pokemon?limit=10000", &data)
	js, _ :=json.Marshal(data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func(u *Handler) GetPokemon(c *gin.Context){
	id := c.Params.ByName("id")
	var data interface{}
	jsonErr := getJson("https://pokeapi.co/api/v2/pokemon/"+id, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c.JSON(500, gin.H{"message": "Error to retrieve pokemons", "error": jsonErr})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"pokemons": data})
	return
}

func(u *HandlerWithoutFramework) GetPokemon(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id :=vars["id"] // the pokemon Id
	var data interface{}
	jsonErr := getJson("https://pokeapi.co/api/v2/pokemon/"+id, &data)
	js, _ :=json.Marshal(data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}