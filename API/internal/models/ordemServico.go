package models

import "gopkg.in/guregu/null.v3"

type OrdemServico struct {
	CodOS       int         `json:"codOS"`
	DataEmissao string      `json:"dataEmissao"`
	CodCli      int         `json:"codCli"`
	NomeCli     string      `json:"nomeCli"`
	Pedido      null.String `json:"pedido"`
	Concluida   bool        `json:"concluida"`
}

func (o OrdemServico) EValida() (bool, string) {
	if o.CodOS < 0 || o.CodOS > 99999999 {
		return false, "Código de ordem de serviço inválido."
	}

	if len(o.DataEmissao) != 10 {
		return false, "Data de emissão deve estar no padrão: YYYY-MM-DD. Ex: 2022-01-30"
	}

	if o.CodCli <= 0 || o.CodCli > 999999 {
		return false, "Código de cliente inválido."
	}

	if o.Pedido.Valid {
		if len(o.Pedido.String) > 255 {
			return false, "A descrição do pedido não pode ultrapassar 255 caracteres."
		}
	}

	return true, ""
}
