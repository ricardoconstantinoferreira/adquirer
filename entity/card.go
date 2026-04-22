package entity

type Card struct {
	Id    int     `json:"id,omitempty"`
	Card  string  `json:"card,omitempty"`
	Flag  string  `json:"flag,omitempty"`
	Cvv   string  `json:"cvv,omitempty"`
	Venc  string  `json:"venc,omitempty"`
	Total float64 `json:"total,omitempty"`
}
