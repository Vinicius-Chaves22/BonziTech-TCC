package models

import "gopkg.in/guregu/null.v3"

type Venda struct {
	CodVenda   int         `json:"codVenda"`
	DataVenda  string      `json:"dataVenda"`
	CodCli     null.Int    `json:"codCli"`
	NomeCli    null.String `json:"nomeCli"`
	CodOS      int         `json:"codOS"`
	ValorTotal float64     `json:"valorTotal"`
	Descricao  null.String `json:"descricao"`
}

func (v Venda) EValida() (bool, string) {
	if v.CodVenda < 0 || v.CodVenda > 99999999 {
		return false, "Código de venda inválido."
	}

	if len(v.DataVenda) != 10 {
		return false, "Data de venda deve estar no padrão: YYYY-MM-DD. Ex: 2022-01-30"
	}

	if v.CodCli.Valid {
		if v.CodCli.Int64 < 0 || v.CodCli.Int64 > 999999 {
			return false, "Código de cliente inválido."
		}
	}

	if v.NomeCli.Valid {
		if v.NomeCli.String == "" || len(v.NomeCli.String) > 30 {
			return false, "Nome do cliente deve ter entre 1 e 30 caracteres."
		}
	}

	if v.CodOS <= 0 || v.CodOS > 99999999 {
		return false, "Código de ordem de serviço inválido."
	}

	if v.ValorTotal <= 0 || v.ValorTotal > 999999999.99 {
		return false, "Valor total de entrada deve ser positivo e menor que 1 bilhão."
	}

	if v.Descricao.Valid {
		if len(v.Descricao.String) > 255 {
			return false, "Descrição da venda não pode ter mais que 255 caracteres."
		}
	}

	v.ValorTotal = float64(int(v.ValorTotal*100) / 100)

	return true, ""
}
