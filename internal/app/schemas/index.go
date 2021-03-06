package schemas

import "time"

type HttpResult struct {
	Result string `json:"result"`
}

type IdResult struct {
	Id int `json:"id"`
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
	// TODO: do auth and get email from redis
	Email string `json:"email"`
}
