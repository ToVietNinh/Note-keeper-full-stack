package presenter

type LoginResponse struct {
	Status  string      `json:"status" example:"200"`
	Message string      `json:"message" example:"SUCCESS"`
	Results *AuthResult `json:"results"`
}

type AuthResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
}
