package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type config struct {
	Next     *string
	Previous *string
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

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
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

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
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
