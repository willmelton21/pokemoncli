package pokeapi

import (
	"io"
	"net/http"
	"log"
	"encoding/json"
	"internal/pokecache"
	"time"
    "fmt"
	)

//api interaction code goes here
type LocationAreasResp struct {
	Count	int
	Next *string
	Previous *string
	Results []struct {
		Name string
		URL string
	}
}

type PokemonAtLocation struct {
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"` // Include if you care about URLs; otherwise omit
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}


type Client struct {
	cache *pokecache.Cache

}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(5 * time.Second), //interval for cache cleanup
	
	}
}
func (c *Client) ExploreLocation(url *string) (*PokemonAtLocation,error) {
	
	PokemonAtLocation := PokemonAtLocation{}
	val, ok := c.cache.Get(*url) 

	if ok {
		err := json.Unmarshal(val,&PokemonAtLocation)
	if err != nil {
		return nil,err
	
	}
	return &PokemonAtLocation,nil
	}
	resp, err := http.Get(*url)
	if err != nil {
		return nil,err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
    if resp.StatusCode == 404 {

		fmt.Printf("Response failed with status code: %d and \nbody: %s\n", resp.StatusCode, body)
        return nil, err
    }
	if resp.StatusCode > 299{

		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", resp.StatusCode, body)
		return nil, err	
	
	}
	if err != nil {
		return nil,err	
	}
	err = json.Unmarshal(body,&PokemonAtLocation)
	if err != nil {
        fmt.Println("couldn't unmarshal")
		return nil,err
	
	}
	c.cache.Add(*url,body)
	return &PokemonAtLocation,nil
}
func (c *Client) GetLocationAreas(nextURL *string) (*LocationAreasResp,error) {
	
	locationResp := LocationAreasResp{}
	val, ok := c.cache.Get(*nextURL) 

	if ok {
		err := json.Unmarshal(val,&locationResp)
	if err != nil {
		return nil,err
	
	}
	return &locationResp,nil
	}
	resp, err := http.Get(*nextURL)
	if err != nil {
		return nil,err
	}
//	fmt.Println("fuffilled get request",resp)
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
    
	if resp.StatusCode > 299{

		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", resp.StatusCode, body)
		return nil, err	
	
	}
//	fmt.Println("body is now:",body)
	if err != nil {
		return nil,err	
	}
	err = json.Unmarshal(body,&locationResp)
	if err != nil {
		return nil,err
	
	}
	c.cache.Add(*nextURL,body)
	return &locationResp,nil
}
