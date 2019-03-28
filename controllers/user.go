package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
)

func (u UserController) GetPing(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"value": "pong"})
	return
}

func (u UserController) GetValueByName(c *gin.Context){
	user := c.Params.ByName("name")
	value, ok := userModel.GetByName(user)
	if ok != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": ok})
		c.Abort()
		return
	}
	if value != nil {
		c.JSON(http.StatusOK, gin.H{"user": value.Name, "value": value.Value})
	} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	return
}

func (u UserController) GetUsers(c *gin.Context){
	
	value, ok := userModel.GetAll()
	if ok != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": ok})
		c.Abort()
		return
	} else{
		c.JSON(http.StatusOK, gin.H{"user": "all", "value": value})
	}
		
	return
}

func (u UserController) GetPokemons(c *gin.Context){
	var urls []string
	mapsUrls := make(map[string]string)
	mapsUrls["pokemon"] = "https://pokeapi.co/api/v2/pokemon?limit=10000"
	mapsUrls["cards"] = "http://www.clashapi.xyz/api/cards"
	mapsUrls["drivers"] = "http://ergast.com/api/f1/drivers?limit=847"
	
	urls = append(urls, "https://pokeapi.co/api/v2/pokemon?limit=10000")
	urls = append(urls, "http://www.clashapi.xyz/api/cards")
	//urls = append(urls, "http://ergast.com/api/f1/drivers?limit=847")

	results := boundedParallelGet(mapsUrls, 2)
	var data interface{} // Pokemons
	var doc []interface{} //cards
	var users MRData //drivers
    var jsonErr error
	for responseIndex := range results {
		binaryResponse, err := ioutil.ReadAll(results[responseIndex].res.Body)
		if results[responseIndex].index == "pokemon"{
		
			log.Println(results[responseIndex].index)
			jsonErr = json.Unmarshal(binaryResponse, &data)
		}
		if results[responseIndex].index == "cards" {
			log.Println(results[responseIndex].index)
			jsonErr = json.Unmarshal(binaryResponse, &doc)
		} else{
			xml.Unmarshal(binaryResponse, &users)
			
			
		}
		if err != nil {
			log.Fatal(err)
		}
		
		
	}
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c.JSON(500, gin.H{"message": "Error to retrieve pokemons", "error": jsonErr})
		c.Abort()
		return
	}
	
	
	/*test, err := getJson("https://pokeapi.co/api/v2/pokemon?limit=10000", data)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"message": "Error to retrieve pokemons", "error": err})
		c.Abort()
		return
	}*/
	c.JSON(http.StatusOK, gin.H{"user": users, "pokemons": data, "cards":doc})
	return
}

func getJson(url string, target interface{})  error{

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	binaryResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil{
		log.Fatal(err)
		return err
	}
	jsonErr := json.Unmarshal(binaryResponse, &target)
	if jsonErr != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
