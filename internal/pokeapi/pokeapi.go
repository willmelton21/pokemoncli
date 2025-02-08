package pokeapi

import (
	"io"
	"net/http"
	"log"
	"encoding/json"
	"internal/pokecache"
	"time"
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

type Client struct {
	cache *pokecache.Cache

}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(5 * time.Second), //interval for cache cleanup
	
	}
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
