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
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type exploreResponse struct {
	PokemonEncounters []struct{ Pokemon struct{ Name string } } `json:"pokemon_encounters"`
}

func commandMap(cfg *config, args []string) error {
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

func commandMapb(cfg *config, args []string) error {
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

func commandExplore(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	target := exploreResponse{}

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

	for _, area := range target.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}

	return nil
}
