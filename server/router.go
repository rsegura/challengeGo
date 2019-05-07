package server

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"challenge/controllers"
	"log"
	//"github.com/gorilla/mux"
)


func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	h, _:=controllers.NewHandler()
	
	//user := new(controllers.UserController)
	// Ping test
	r.GET("/ping", h.GetPing)

	// Get user value
	//r.GET("/user/:name", h.GetValueByName)

	r.GET("/pokemons", h.GetAllPokemons)
	r.GET("/clash", h.GetAllClash)
	r.GET("/clash/:id", h.GetClash)
	r.GET("/pokemon/:id", h.GetPokemon)
	r.GET("/elements",h.GetAllValues)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	/*authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})*/

	return r
}

func(a *App) SetupRouterWithoutFramework(){
	log.Println("On SetupRouterWithoutFramework")
	h, _:=controllers.NewHandlerWithoutFramework()
	a.Router.HandleFunc("/ping", h.GetPing)
	a.Router.HandleFunc("/pokemons", h.GetAllPokemons)
	a.Router.HandleFunc("/pokemon/{id}", h.GetPokemon)
	a.Router.HandleFunc("/all", h.GetAllValues)
}