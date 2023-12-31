package entity

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Balance  Balance
}

type Balance struct {
	Money int `json:"current" db:"current"`
	Bonus int `json:"withdrawn" db:"withdraw"`
}
