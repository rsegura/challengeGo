package server

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

type App struct{
	Router *mux.Router
}

func(a *App) Init(){
	log.Println("Init Server")
	a.Router = mux.NewRouter()
	a.SetupRouterWithoutFramework()
	amw := authenticationMiddleware{}
	amw.tokenUsers = make(map[string]string)
	amw.Populate()
	//a.Router.Use(amw.Middleware)
	a.Router.Use(loggingMiddleware)
}

func (a *App) Run(addr string){
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}



// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Session-Token")

        if user, found := amw.tokenUsers[token]; found {
        	// We found the token in our map
        	log.Printf("Authenticated user %s\n", user)
        	// Pass down the request to the next middleware (or final handler)
        	next.ServeHTTP(w, r)
        } else {
        	// Write an error and stop the handler chain
        	http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}