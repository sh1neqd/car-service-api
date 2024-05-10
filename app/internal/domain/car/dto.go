package car

type CreateCarDTO struct {
	RegNums []string `json:"reg_nums"`
}

type CarDTO struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
}

type UpdateCarDTO struct {
	RegNum *string `json:"regNum"`
	Mark   *string `json:"mark"`
	Model  *string `json:"model"`
}
