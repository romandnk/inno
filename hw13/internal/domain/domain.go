package domain

type Ing struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Recipe struct {
	ID          string `json:"id"`
	UserID      string
	Name        string `json:"name"`
	Ingredients []Ing  `json:"ingredients"`
	Temperature int    `json:"temperature"`
}
