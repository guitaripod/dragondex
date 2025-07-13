package models

type Character struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Description string    `json:"description"`
	PowerLevel  int       `json:"power_level"`
	Abilities   []Ability `json:"abilities"`
	Episodes    []Episode `json:"episodes"`
	ImageURL    string    `json:"image_url"`
}
