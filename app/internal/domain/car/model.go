package car

type Car struct {
	ID     int    `json:"id" db:"id"`
	RegNum string `json:"regNum" db:"reg_num"`
	Mark   string `json:"mark" db:"mark"`
	Model  string `json:"model" db:"model"`
}
