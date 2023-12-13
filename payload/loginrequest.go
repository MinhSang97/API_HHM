package payload

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Error  string      `json:"error,omitempty"`
}
