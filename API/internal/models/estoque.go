package models

type Estoque struct {
	CodEstq    int     `json:"codEstq"`
	CodComp    int     `json:"codComp"`
	QuantMin   int     `json:"min"`
	QuantMax   int     `json:"max"`
	QuantAtual float64 `json:"quantidade"`
}

func (e Estoque) EValido() (bool, string) {
	if e.CodEstq < 0 || e.CodEstq > 99999999 {
		return false, "Código de estoque inválido."
	}

	if e.CodComp <= 0 || e.CodComp > 99999999 {
		return false, "Código de componente inválido."
	}

	if e.QuantMin < 0 || e.QuantMin > 999 {
		return false, "Quantidade mínima deve ser inferior a 1.000."
	}

	if e.QuantMax <= 0 || e.QuantMax > 9999 {
		return false, "Quantidade máxima deve ser maior que 0 e inferior a 10.000."
	}

	if e.QuantMin > e.QuantMax {
		return false, "Quantidade mínima permitida não pode ser maior que a máxima."
	}

	if e.QuantAtual < 0 {
		return false, "Quantidade de componentes inválida."
	}

	return true, ""
}
