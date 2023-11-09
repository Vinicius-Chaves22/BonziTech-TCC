package models

import "strings"

type Usuario struct {
	CodUsuario int    `json:"codUsuario"`
	Permissoes string `json:"permissoes"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Senha      string `json:"senha"`
	Ativo      bool   `json:"ativo"`
}

type UsuarioResponse struct {
	CodUsuario int    `json:"codUsuario"`
	Permissoes string `json:"permissoes"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Senha      string `json:"senha"`
	Ativo      bool   `json:"ativo"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func (u Usuario) EValido() (bool, string) {
	if u.CodUsuario < 0 || u.CodUsuario > 999999 {
		return false, "Código de usuário inválido."
	}

	if u.Permissoes == "" || len(u.Permissoes) > 20 {
		return false, "Permissão de usuário deve conter de 1 a 20 caracteres."
	}

	if u.Permissoes != "ADM" && u.Permissoes != "Administrador" &&
		u.Permissoes != "Leitura" {
		return false, "Usuário pode ter apenas permissões de Administrador e Leitura."
	}

	if u.Nome == "" || len(u.Nome) > 30 {
		return false, "Nome de usuário deve conter de 1 a 30 caracteres."
	}

	if len(u.Email) < 5 || len(u.Email) > 70 {
		return false, "Email deve conter de 5 a 70 caracteres."
	}

	if !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
		return false, "Email deve conter um @ (arroba), e um . (ponto)."
	}

	if len(u.Senha) < 8 {
		return false, "Senha deve conter pelo menos 8 caracteres."
	}

	return true, ""
}
