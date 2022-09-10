package todo

type Request struct {
	Id           int      `json:"id" db:"id"`
	Purpose      string   `json:"purpose" db:"purpose"`
	Description  string   `json:"description" db:"description"`
	Status       string   `json:"status" db:"status"`
	Amount       float64  `json:"amount" db:"amount"`
	Images       []string `json:"images" db:"images"`
	UserId       int      `json:"user_id" db:"user_id"`
	BookkeeperId int      `json:"bookkeeper_id" db:"bookkeeper_id"`
}
