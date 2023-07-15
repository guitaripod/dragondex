package models

type Ability struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PowerRequired int    `json:"power_required"`
}
