package main

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

type AreaData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} 
}

func helperAreaGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func commandMap(conf *config, _ string) error {
	var data []byte
	data, ok := conf.cache.Get(conf.areaNext)
	if !ok {
		var err error
		data, err = helperAreaGet(conf.areaNext)
		if err != nil {
			return err
		}
		conf.cache.Add(conf.areaNext, data)
	}
	areaDat := AreaData{}
	if err := json.Unmarshal(data, &areaDat); err != nil {
		return  err 
	}
	conf.areaNext = areaDat.Next
	conf.areaPrev = areaDat.Previous
	for _, name := range areaDat.Results {
		fmt.Printf("%s\n", name.Name)
	}
	return nil
}

func commandMapb(conf *config, _ string) error {
	prev, ok := conf.areaPrev.(string)
	if !ok {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	var data []byte
	data, ok = conf.cache.Get(prev)
	if !ok {
		var err error
		data, err = helperAreaGet(prev)
		if err != nil {
			return err
		}
		conf.cache.Add(prev, data)
	}
	areaDat := AreaData{}
	if err := json.Unmarshal(data, &areaDat); err != nil {
		return  err 
	}
	conf.areaNext = areaDat.Next
	conf.areaPrev = areaDat.Previous
	for _, name := range areaDat.Results {
		fmt.Printf("%s\n", name.Name)
	}
	return nil
}
