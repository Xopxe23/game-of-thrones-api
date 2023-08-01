package gameofthrones

import "fmt"

type Person struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type House struct {
	Slug    string   `json:"slug"`
	Name    string   `json:"name"`
	Members []Person `json:"members"`
}

func (p Person) Info() string {
	return fmt.Sprintf("['%s' %s]", p.Slug, p.Name)
}

func (h House) Info() string {
	var members string
	for _, member := range h.Members {
		members += member.Info()
	}
	return fmt.Sprintf("[%s] %s members: %s", h.Slug, h.Name, members)
}
