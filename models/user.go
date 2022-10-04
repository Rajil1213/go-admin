package models

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  []byte `json:"-"`
}
