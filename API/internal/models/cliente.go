package models

import (
	"strings"

	"gopkg.in/guregu/null.v3"
)

type Cliente struct {
	CodCli      int         `json:"codCli"`
	NomeEmpresa string      `json:"nomeEmpresa"`
	NomeCli     string      `json:"nome"`
	Tipo        null.String `json:"tipo"`
	DiaReg      string      `json:"diaReg"`
	Endereco    null.String `json:"endereco"`
	Bairro      null.String `json:"bairro"`
	Cidade      string      `json:"cidade"`
	Estado      string      `json:"estado"`
	CEP         null.String `json:"cep"`
	Email       null.String `json:"email"`
}

func (c Cliente) EValido() (bool, string) {
	if c.CodCli < 0 || c.CodCli > 999999 {
		return false, "Código de cliente inválido."
	}

	if c.NomeEmpresa == "" || len(c.NomeEmpresa) > 70 {
		return false, "Nome de empresa deve conter de 1 a 70 caracteres."
	}

	if c.NomeCli == "" || len(c.NomeCli) > 30 {
		return false, "Nome de cliente deve conter de 1 a 30 caracteres."
	}

	if c.Tipo.Valid {
		if len(c.Tipo.String) > 32 {
			return false, "Tipo de cliente deve conter até 32 caracteres."
		}
	}

	if c.Endereco.Valid {
		if len(c.Endereco.String) > 128 {
			return false, "Endereço deve conter até 128 caracteres."
		}
	}

	if c.Bairro.Valid {
		if len(c.Bairro.String) > 30 {
			return false, "Bairro deve conter até 30 caracteres."
		}
	}

	if c.Cidade == "" || len(c.Cidade) > 30 {
		return false, "Cidade deve conter até 30 caracteres."
	}

	if len(c.Estado) != 2 {
		return false, "Estado deve ser a sigla do estado. Ex: SP, RJ, PB, ..."
	}

	if c.CEP.Valid {
		if len(c.CEP.String) != 9 {
			return false, "CEP deve conter 9 caracteres. Ex: 10000-000"
		}
	}

	if c.Email.Valid {
		if !strings.Contains(c.Email.String, "@") ||
			!strings.Contains(c.Email.String, ".") {
			return false, "Email deve conter um @ (arroba), e um . (ponto)."
		}
	}

	return true, ""
}
