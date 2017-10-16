package models

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
	DeadlineDate string `json:"deadline_date"`
}