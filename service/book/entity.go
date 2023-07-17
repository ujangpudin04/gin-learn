package book

import "encoding/json"

type Book struct {
	ID          json.Number `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
	Qty         string `json:"qty"`
}
