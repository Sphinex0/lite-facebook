package models

type User struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	DateBirth  string    `json:"date_birth"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Image      string `json:"image"`
	AboutMe    string `json:"about"`
	Privacy    string `json:"privacy"`
	CreatedAt  string `json:"created_at"`
}

type UserInfo struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Image      string `json:"image"`
}