package models

type Image struct {
	Link   string `json:"link"`
	Camera Camera `json:"camera"`
	Rover  Rover  `json:"rover"`
	Date   string `json:"formatted_date"`
}
