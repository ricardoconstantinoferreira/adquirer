package entity

type AuthorizationResponse struct {
	Token    	string      `json:"token"`
	Code   		string      `json:"code"`
	Message 	string 		`json:"message"`
}