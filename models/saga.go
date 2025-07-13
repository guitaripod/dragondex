package models

import "time"

type Saga struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Episodes    []Episode   `json:"episodes"`
	Characters  []Character `json:"characters"`
}
