package entity

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Balance  Balance
}

type Balance struct {
	Money float64 `json:"current" db:"current"`
	Bonus int     `json:"withdraw" db:"withdraw"`
}
