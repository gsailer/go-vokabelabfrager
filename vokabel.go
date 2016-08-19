package main

type Vokabelset struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Choices  []string `json:"choices"`
}
type Sets []Vokabelset
