package book

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
	Qty         string `json:"qty"`
	IdPerson    int    `db:"id_person" json:"id_person"`
	UpdatedAt   string `json:"updated_at"`
}
