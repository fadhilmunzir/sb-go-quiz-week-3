package user

import "time"

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Created_at  time.Time `json:"created_at"`
	Created_by  string    `json:"created_by"`
	Modified_at time.Time `json:"modified_at"`
}
