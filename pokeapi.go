package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/TTTV273/Pokedex/internal/pokecache"
)

type Pokemon struct {
	Name           string
	BaseExperience int `json:"base_experience"`
	Height         int
	Weight         int
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		}
	}
	Types []struct {
		Type struct {
			Name string `json:"name"`
		}
	}
}

type config struct {
	Next     *string
	Previous *string
	Cache    pokecache.Cache
	Pokedex  map[string]Pokemon
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

func commandCatch(cfg *config, args []string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	target := Pokemon{}

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

	threshold := 50
	randomNum := rand.Intn(target.BaseExperience)
	if randomNum < threshold {
		fmt.Printf("%v was caught!\n", args[0])
		cfg.Pokedex[args[0]] = target
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}

	return nil
}

func commandInspect(cfg *config, args []string) error {
	if pokemon, ok := cfg.Pokedex[args[0]]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, name := range pokemon.Types {
			fmt.Printf("  -%v\n", name.Type.Name)
		}
	} else {
		fmt.Printf("you have not caught that pokemon\n")
	}

	return nil
}
