package models

type Profile struct {
	UserInfo `json:"user_info"`
	Followers int `json:"followers"`
	Followings int `json:"followings"`
	Action string  `json:"action"`
}
