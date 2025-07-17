package models

import "time"

type Tasks struct {
	ID          string    `json: "id"`
	Title       string    `json: "title"`
	Description string    `json: "description"`
	DueDate     time.Time `json: "duedate"`
	Status 		string	  `json: "status"`
}