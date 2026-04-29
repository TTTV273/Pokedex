package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TTTV273/Pokedex/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	Cache    pokecache.Cache
}

type locationAreaResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct{ Name string }
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	target := locationAreaResponse{}

	if cfg.Next != nil {
		url = *cfg.Next
	}

	var body []byte
	var err error
	var res *http.Response
	if cachedBody, ok := cfg.Cache.Get(url); ok {
		body = cachedBody
	} else {
		res, err = http.Get(url)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}

		cfg.Cache.Add(url, body)
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, area := range target.Results {
		fmt.Println(area.Name)
	}
	cfg.Next = target.Next
	cfg.Previous = target.Previous

	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	target := locationAreaResponse{}

	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url = *cfg.Previous

	var body []byte
	var err error
	var res *http.Response
	if cachedBody, ok := cfg.Cache.Get(url); ok {
		body = cachedBody
	} else {
		res, err = http.Get(url)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}

		cfg.Cache.Add(url, body)
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, area := range target.Results {
		fmt.Println(area.Name)
	}
	cfg.Next = target.Next
	cfg.Previous = target.Previous

	return nil
}
