package models

import "gopkg.in/guregu/null.v3"

type Sessao struct {
	CodSessao  int         `json:"codSessao"`
	CodUsuario int         `json:"codUsuario"`
	Entrada    string      `json:"entrada"`
	Saida      null.String `json:"saida"`
}

type SessaoResponse struct {
	CodSessao  int         `json:"codSessao"`
	CodUsuario int         `json:"codUsuario"`
	Entrada    string      `json:"entrada"`
	Saida      null.String `json:"saida"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
}
