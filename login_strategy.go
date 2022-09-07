package ews

import "net/http"

type LoginStrategy interface {
	SetLoginHeaders(req *http.Request)
}

type PlainLogin struct {
	Username string
	Password string
}

func (p PlainLogin) SetLoginHeaders(req *http.Request) {
	req.SetBasicAuth(p.Username, p.Password)
}

type XOAuthLogin struct{ Token string }

func (p XOAuthLogin) SetLoginHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+p.Token)
}
