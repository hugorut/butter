package httpio

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorOutput struct {
	Error string `json:"error"`
}
