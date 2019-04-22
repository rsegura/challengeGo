package controllers

import(
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Clash struct {
	Id string `json:"_id"`
	IdName string
	Name string
	Image string
}


func(u *Handler) GetAllClash(c *gin.Context){
	var data []interface{}
	jsonErr := getJson("http://www.clashapi.xyz/api/cards", &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c.JSON(500, gin.H{"message": "Error to retrieve clash elements", "error": jsonErr})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"clash": data})
	return
}

func(u *Handler) GetClash(c *gin.Context){
	id := c.Params.ByName("id")
	var data Clash
	jsonErr := getJson("http://www.clashapi.xyz/api/cards/"+id, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		c.JSON(500, gin.H{"message": "Error to retrieve clash element", "error": jsonErr})
		c.Abort()
		return
	}
	data.Image="/images/cards/"+data.IdName+".png"
	c.JSON(http.StatusOK, gin.H{"clash": data})
	return

}