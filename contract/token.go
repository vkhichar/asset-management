package contract

type TokenClaims struct {
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
