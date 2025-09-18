package user

type TokenResponse struct {
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
