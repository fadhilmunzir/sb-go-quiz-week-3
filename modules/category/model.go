package category

import "time"

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Created_at  time.Time `json:"created_at"`
	Created_by  string    `json:"created_by"`
	Modified_at time.Time `json:"modified_at"`
	Modified_by string    `json:"modified_by"`
}
