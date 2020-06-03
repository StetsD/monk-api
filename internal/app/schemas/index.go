package schemas

import "time"

type HttpResult struct {
	Result string `json:"result"`
}

type RegistrationBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EventBody struct {
	Title       string    `json:"title"`
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	Description string    `json:"description"`
	UserId      int64     `json:"userId"`
}
