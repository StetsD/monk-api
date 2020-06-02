package schemas

type HttpResult struct {
	Result string `json:"result"`
}

type RegistrationBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
