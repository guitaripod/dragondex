package models

type Episode struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	AirDate     time.Time   `json:"air_date"`
	Description string      `json:"description"`
	Saga        Saga        `json:"saga"`
	Characters  []Character `json:"characters"`
}
