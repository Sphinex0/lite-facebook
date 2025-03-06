package models

type User struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Dob        string `json:"dob"`
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Uuid       string `json:"uuid"`
	Uuid_exp   int    `json:"uuid_exp"`
	Image      string `json:"image"`
	AboutMe    string `json:"aboutMe"`
}

type UserInfo struct {
	Nickname   string `json:"nickname"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Image      string `json:"image"`
}
