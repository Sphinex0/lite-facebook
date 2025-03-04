package models

type User struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Age        string `json:"age"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Uuid       string `json:"uuid"`
	Uuid_exp   int    `json:"uuid_exp"`
	Image      string `json:"image"`
}

type UserInfo struct {
	Nickname   string `json:"nickname"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Image      string `json:"image"`
}
