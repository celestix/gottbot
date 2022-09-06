package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Schemas struct {
	Properties Properties `json:"properties,omitempty"`
	Required   []string   `json:"required,omitempty"`
}
type Properties struct {
	Description any `json:"description,omitempty"`
	Type        any `json:"type"`
	Format      any `json:"format,omitempty"`
	Nullable    any `json:"nullable,omitempty"`
}

type Components struct {
	Schemas map[string]Schemas `json:"schemas"`
}

type API struct {
	Components Components `json:"components"`
	// Paths any `json:"paths"`
}

func main() {
	b, err := ioutil.ReadFile("schema.json")
	if err != nil {
		panic(err.Error())
	}
	var api API
	if err = json.Unmarshal(b, &api); err != nil {
		panic(err.Error())
	}
	fmt.Println(api)
}
