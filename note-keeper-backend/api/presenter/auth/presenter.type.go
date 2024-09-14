package presenter

type BasicResponse struct {
	Status  string      `json:"status" example:"200"`
	Message string      `json:"message" example:"SUCCESS"`
	Results interface{} `json:"results"`
}

type Error400Response struct {
	Status  string      `json:"status" example:"400"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}
type Error404Response struct {
	Status  string      `json:"status" example:"500"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}
type Error500Response struct {
	Status  string      `json:"status" example:"500"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}

type AuthResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results *AuthResult `json:"results"`
}

type AuthResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
}

type TokenResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Results *TokenResult `json:"results"`
}

type TokenResult struct {
	Token string `json:"token"`
}
