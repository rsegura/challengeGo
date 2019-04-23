package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"io/ioutil"
	"sort"
	"encoding/json"
	"encoding/xml"
)

type result struct {
	index string
	res   http.Response
	err   error
}

type MRData struct{
	XMLName xml.Name `xml:"MRData" json:"-"`
	Series string `xml:"series,attr" json:"series"`
	Url string `xml:"url,attr" json:"url"`
	Limit string `xml:"limit,attr" json:"limit"`
	Total string `xml:"total,attr" json:"total"`
	DriversList []Driver `xml:"DriverTable>Driver" json:"drivers"`
}

type Driver struct{
	XMLName xml.Name `xml:"Driver" json:"-"`
	DriverId string `xml:"driverId,attr" json:"driverId"`
	Url string	`xml:"url,attr" json:"url"`
	GivenName string `xml:"GivenName" json:"givenName"`
	FamilyName string `xml:"FamilyName" json:"familyName"`
	DateOfBirth string `xml:"DateOfBirth" json:"dateOfBirth"`
	Nationality string `xml:"Nationality" json:"nationality"`
}


type UserController struct{

}

type pokemonRequest struct{
	Count int `json:"count"`
	Next int `json:"next"`
	Previous int `json:"previous"`	
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type cardsRequest struct{
	Arena int `json:"arena"`
	CopyId int `json:"copyId"`
	Description string `json:"description"`
	ElixirCost int `json:"elixirCost"`
	IdName string `json:"idName"`
	Name string `json:"name"`
	Order int `json:"order"`
	Rarity string `json:"rarity"`
	Type string `json:"type"`
}


func (u *Handler) GetAllValues(c *gin.Context){
	
	mapsUrls := make(map[string]string)
	mapsUrls["pokemon"] = "https://pokeapi.co/api/v2/pokemon?limit=10000"
	mapsUrls["cards"] = "http://www.clashapi.xyz/api/cards"
	mapsUrls["drivers"] = "http://ergast.com/api/f1/drivers?limit=847"
	
	results := boundedParallelGet(mapsUrls, 3)
	var data pokemonRequest // Pokemons try to change this to struct in order to use real data types
	var doc []cardsRequest //cards   same
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

	c.JSON(http.StatusOK, gin.H{"user": users, "pokemons": data, "cards":doc})
	return
}

func boundedParallelGet(urls map[string]string, concurrencyLimit int) []result {

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the http request results
	resultsChan := make(chan *result)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keen an index and loop through every url we will send a request to
	for key, url := range urls {

		// start a go routine with the index and url in a closure
		go func(key string, url string) {

			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			// send the request and put the response in a result struct
			// along with the index so we can sort them later along with
			// any error that might have occoured
			res, err := http.Get(url)
			result := &result{key, *res, err}

			// now we can send the result struct through the resultsChan
			resultsChan <- result

			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(key, url)
	}

	// make a slice to hold the results we're expecting
	var results []result

	// start listening for any results over the resultsChan
	// once we get a result append it to the result slice
	for {
		result := <-resultsChan
		results = append(results, *result)

		// if we've reached the expected amount of urls then stop
		if len(results) == len(urls) {
			break
		}
	}

	// let's sort these results real quick
	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	// now we're done we return the results
	return results
}