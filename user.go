package todo

type User struct {
	ID int `json:"-" db:"id"`
	Name string `json:"name" blinding:"required"`
	Userame string `json:"username" blinding:"required"`
	Password string `json:"password" blinding:"required"`
}