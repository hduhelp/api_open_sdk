package auth

import (
	"encoding/json"
	"github.com/hduhelp/api_open_sdk/request"
)

type TokenValid struct {
	AccessToken string `json:"accessToken"`
	ClientID    string `json:"clientID"`
	ExpiredTime int    `json:"expiredTime"`
	GrantType   string `json:"grantType"`
	IsValid     int    `json:"IsValid"`
	StaffID     string `json:"staffId"`
	TokenType   int    `json:"tokenType"`
	Uid         int    `json:"uid"`
}

func (v *TokenValid) Unmarshal(i []byte) error {
	return json.Unmarshal(i, &v)
}

func GetTokenValidDetail(token string) *TokenValid {
	data := new(TokenValid)
	request.New().SetToken(token).
		EndData(data)
	return data
}
