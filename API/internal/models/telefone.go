package models

import "gopkg.in/guregu/null.v3"

type Telefone struct {
	CodTel      int         `json:"codTel"`
	CodCli      int         `json:"codCli"`
	Telefone    string      `json:"telefone"`
	NomeTel     string      `json:"nomeTel"`
	TipoContato null.String `json:"tipoContato"`
	TipoCli     null.String `json:"tipoCli"`
}

func (t Telefone) EValido() (bool, string) {
	if t.CodTel < 0 || t.CodTel > 99999999 {
		return false, "Código de telefone inválido."
	}

	if t.CodCli <= 0 || t.CodCli > 999999 {
		return false, "Código de cliente inválido."
	}

	if len(t.Telefone) < 8 || len(t.Telefone) > 19 {
		return false, "Um telefone deve ter de 8 a 19 caracteres."
	}

	if t.NomeTel == "" || len(t.NomeTel) > 45 {
		return false, "Nome do telefone deve conter de 1 a 45 caracteres."
	}

	if t.TipoContato.Valid {
		if len(t.TipoContato.String) > 30 {
			return false, "Tipo de contato deve conter até 30 caracteres."
		}
	}

	if t.TipoCli.Valid {
		if len(t.TipoCli.String) > 30 {
			return false, "Tipo de cliente deve conter até 30 caracteres."
		}
	}

	return true, ""
}
