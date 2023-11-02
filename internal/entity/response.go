package entity

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    *User  `json:"data"`
}

type UserResponse struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    []ListUser `json:"data"`
}

type UserAffiliatedResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    *ListUser `json:"data"`
}
