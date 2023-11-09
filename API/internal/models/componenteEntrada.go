package models

type ComponenteEntrada struct {
	CodCompEntd int     `json:"codCompEntd"`
	CodEntd     int     `json:"codEntd"`
	CodComp     int     `json:"codComp"`
	Quantidade  float64 `json:"quantidade"`
	ValorUnit   float64 `json:"valorUnit"`
}

func (c ComponenteEntrada) EValido() (bool, string) {
	if c.CodCompEntd < 0 || c.CodCompEntd > 99999999 {
		return false, "Código de componente de entrada inválido."
	}

	if c.CodEntd <= 0 || c.CodEntd > 99999999 {
		return false, "Código de entrada inválido."
	}

	if c.CodComp <= 0 || c.CodComp > 99999999 {
		return false, "Código de componente inválido."
	}

	if c.Quantidade <= 0 {
		return false, "A quantidade de um componente em entrada deve ser maior que 0."
	}

	if c.ValorUnit <= 0 || c.ValorUnit > 9999999.99 {
		return false, "Valor de componente deve estar entre 0,01 e 10.000.000 R$."
	}

	c.ValorUnit = float64(int(c.ValorUnit*100) / 100)

	return true, ""
}
