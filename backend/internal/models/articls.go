package models

type Articls struct {
	ID string `json:"id"`
	User_Id string `json:"user_id"`
	Content string `json:"content"`
	Privacy string `json:"privacy"`
	Created_at string `json:"created_at"`
	Modified_at string `json:"modified_at"`
	Image string `json:"image"`
	Parent string `json:"parent"`
	Group_id string `json:"group_id"`
}
