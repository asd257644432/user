package model

type User struct {
	Id       int64  `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Passport string `json:"passport"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
