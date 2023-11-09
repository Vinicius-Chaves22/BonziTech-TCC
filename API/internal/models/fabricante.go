package models

import "gopkg.in/guregu/null.v3"

type Fabricante struct {
	CodFab      int         `json:"codFab"`
	Nome        string      `json:"nome"`
	NomeContato null.String `json:"nomeContato"`
	RazaoSocial null.String `json:"razaoSocial"`
	Telefone    null.String `json:"telefone"`
	Celular     null.String `json:"celular"`
	Fax         null.String `json:"fax"`
	Endereco    null.String `json:"endereco"`
	Cidade      null.String `json:"cidade"`
	Estado      null.String `json:"estado"`
	CEP         null.String `json:"cep"`
}

func (f Fabricante) EValido() (bool, string) {
	if f.CodFab < 0 || f.CodFab > 999999 {
		return false, "Código de fabricante inválido."
	}

	if f.Nome == "" || len(f.Nome) > 45 {
		return false, "Nome de fabricante deve ter de 1 a 45 caracteres."
	}

	if f.NomeContato.Valid {
		if len(f.NomeContato.String) > 50 {
			return false, "Nome de contato deve ter no máximo 50 caracteres."
		}
	}

	if f.RazaoSocial.Valid {
		if len(f.RazaoSocial.String) > 60 {
			return false, "Razão social de fabricante deve conter até 60 caracteres."
		}
	}

	if f.Telefone.Valid {
		if len(f.Telefone.String) > 19 {
			return false, "Telefone deve conter até 19 caracteres."
		}
	}

	if f.Celular.Valid {
		if len(f.Celular.String) > 19 {
			return false, "Celular deve conter até 19 caracteres."
		}
	}

	if f.Fax.Valid {
		if len(f.Fax.String) > 19 {
			return false, "Fax deve conter até 19 caracteres."
		}
	}

	if f.Endereco.Valid {
		if len(f.Endereco.String) > 128 {
			return false, "Endereço deve conter até 128 caracteres."
		}
	}

	if f.Cidade.Valid {
		if len(f.Cidade.String) > 30 {
			return false, "Cidade deve conter até 30 caracteres."
		}
	}

	if f.Estado.Valid {
		if len(f.Estado.String) != 2 {
			return false, "Estado deve ser a sigla do estado. Ex: SP, RJ, PB, ..."
		}
	}

	return true, ""
}
