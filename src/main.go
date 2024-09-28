package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

type NpmRegistryResponse struct {
	Name     string `json:"name"`
	Versions map[string]struct {
		Repository struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:".repository"`
		HomePage string `json:"homepage"`
	} `json:"versions"`
}

const SEARCH_URL string = "https://registry.npmjs.org/"

func main() {
	// const packName string = "mongoose"
	if len(os.Args) < 2 {
		panic("No package name provided as argument")
	}
	var packName string = os.Args[1]
	fmt.Println("Searching for", packName)
	reqUrl := SEARCH_URL + packName
	fmt.Println(reqUrl)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header = http.Header{
		"Accept": {"application/json"},
	}
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var body NpmRegistryResponse
	data, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		panic(err)
	}
	var keys []string
	for k := range body.Versions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	latest := keys[len(keys)-1]
	fmt.Println(latest)
}
