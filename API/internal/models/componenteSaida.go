package models

type ComponenteSaida struct {
	CodCompVenda int     `json:"codCompVenda"`
	CodVenda     int     `json:"codVenda"`
	CodComp      int     `json:"codComp"`
	Quantidade   float64 `json:"quantidade"`
	ValorUnit    float64 `json:"valorUnit"`
}

func (c ComponenteSaida) EValido() (bool, string) {
	if c.CodCompVenda < 0 || c.CodCompVenda > 99999999 {
		return false, "Código de componente de venda inválido."
	}

	if c.CodVenda <= 0 || c.CodVenda > 99999999 {
		return false, "Código de saída inválido."
	}

	if c.CodComp <= 0 || c.CodComp > 99999999 {
		return false, "Código de componente inválido."
	}

	if c.Quantidade <= 0 {
		return false, "A quantidade de um componente em saída deve ser maior que 0."
	}

	if c.ValorUnit <= 0 || c.ValorUnit > 9999999.99 {
		return false, "Valor de componente deve estar entre 0,01 e 10.000.000 R$."
	}

	c.ValorUnit = float64(int(c.ValorUnit*100) / 100)

	return true, ""
}
