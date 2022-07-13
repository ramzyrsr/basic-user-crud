package model

type User struct {
	UserId   uint32 `json:"userId"`
	Email    string `json:"email"  validate:"required"`
	Address  string `json:"address"`
	Password string `json:"password"  validate:"required"`
}

func (User) UserTable() string {
	return "users"
}
