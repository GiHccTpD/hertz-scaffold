package model

type Response struct {
	Code    string      `json:"code" example:"0000"`
	Message string      `json:"message" example:"OK"`
	Data    interface{} `json:"data"`
}

type IdToken struct {
	IdToken string `json:"id_token" example:"eyJhbGciOiJSUzI1NiIxxxxxxx"`
}

type OKResponse struct {
	Code    string  `json:"code" example:"0000"`
	Message string  `json:"message" example:"OK"`
	Data    IdToken `json:"data"`
}

type TokenErrorResponse struct {
	Code    string      `json:"code" example:"0705"`
	Message string      `json:"message" example:"token error"`
	Data    interface{} `json:"data"`
}

type ServerErrorResponse struct {
	Code    string      `json:"code" example:"0011"`
	Message string      `json:"message" example:"server internal error"`
	Data    interface{} `json:"data"`
}
