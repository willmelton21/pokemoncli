package pokeapi

import (
	"io"
	"net/http"
	"log"
	"encoding/json"
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

func GetLocationAreas(nextURL *string) (*LocationAreasResp,error) {

	locationResp := LocationAreasResp{}
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
	return &locationResp,nil
}
