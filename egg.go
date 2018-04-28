package main

type Egg struct {
	Nick        string `json: "nick"`
	HasData     bool   `json:"hasData"`
	Description string `json:"description"`
}
