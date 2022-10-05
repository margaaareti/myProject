package models

type User struct {
	Id         uint64 `form:"id" db:"id"`
	Name       string `form:"name"`
	Surname    string `form:"surname"`
	Patronymic string `form:"patronymic"`
	Username   string `form:"username"`
	Password   string `form:"password"`
	Email      string `form:"email"`
}
