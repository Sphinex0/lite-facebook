package models

type Error struct {
	Err  error `json:"err"`
	Code int   `json:"code"`
}
