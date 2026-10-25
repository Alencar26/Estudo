package entities

type Ingredients []string

type Coffee struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Ingredients Ingredients `json:"ingredients"`
	Image       string      `json:"image"`
}
