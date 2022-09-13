package todo

type Deal struct {
	Id           int      `json:"id" db:"id"`
	Purpose      string   `json:"purpose" db:"purpose"`
	Description  string   `json:"description" db:"description"`
	Status       string   `json:"status" db:"status"`
	Amount       float64  `json:"amount" db:"amount"`
	Images       []string `json:"images" db:"images"`
	UserId       int      `json:"user_id" db:"user_id"`
	BookkeeperId int      `json:"bookkeeper_id" db:"bookkeeper_id"`
}

type AllNewDeals struct {
	Id           int     `json:"id" db:"id"`
	Purpose      string  `json:"purpose" db:"purpose"`
	Status       string  `json:"status" db:"status"`
	Amount       float64 `json:"amount" db:"amount"`
	FullName     string  `json:"full_name" db:"full_name"`
	BookkeeperId int     `json:"bookkeeper_id" db:"bookkeeper_id"`
	CreatedAt    string  `json:"created_at" db:"created_at"`
}

type AllOwnDeal struct {
	Status string `json:"status" db:"status"`
}

type OneDeal struct {
	Id          int      `json:"id" db:"id"`
	Purpose     string   `json:"purpose" db:"purpose"`
	Description string   `json:"description" db:"description"`
	Status      string   `json:"status" db:"status"`
	Amount      float64  `json:"amount" db:"amount"`
	Images      []string `json:"images" db:"images"`
	UserId      int      `json:"user_id" db:"user_id"`
	FullName    string   `json:"full_name" db:"full_name"`
	CreatedAt   string   `json:"created_at" db:"created_at"`
}

type UpdateDealStatus struct {
	Id     int    `json:"deal_id" db:"id"`
	Status string `json:"status" db:"status"`
}

type UpdateDealBookkeeperId struct {
	RequestId int `json:"request_id"`
}
