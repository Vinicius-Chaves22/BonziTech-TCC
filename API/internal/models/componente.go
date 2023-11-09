package models

import "gopkg.in/guregu/null.v3"

type Componente struct {
	CodComp          int         `json:"codComp"`
	CodPeca          string      `json:"codPeca"`
	Especificacao    string      `json:"especificacao"`
	CodCat           int         `json:"codCat"`
	CodSubcat        null.Int    `json:"codSubcat"`
	DiamInterno      null.String `json:"diamInterno"`
	DiamExterno      null.Float  `json:"diamExterno"`
	DiamNominal      null.String `json:"diamNominal"`
	MedidaD          null.Int    `json:"medidaD"`
	Costura          null.Bool   `json:"costura"`
	PrensadoReusavel null.String `json:"prensadoReusavel"`
	Mangueira        null.String `json:"mangueira"`
	Material         null.String `json:"material"`
	Norma            null.String `json:"norma"`
	Bitola           null.Int    `json:"bitola"`
	ValorEntrada     float64     `json:"valorEntrada"`
	ValorSaida       float64     `json:"valorVenda"`
}

func (c Componente) EValido() (bool, string) {
	if c.CodComp < 0 || c.CodComp > 99999999 {
		return false, "Código de componente inválido."
	}

	if c.CodPeca == "" || len(c.CodPeca) > 30 {
		return false, "Código de peça deve conter de 1 a 30 caracteres."
	}

	if c.Especificacao == "" || len(c.Especificacao) > 100 {
		return false, "Especificação de componente deve conter de 1 a 100 caracteres."
	}

	if c.CodCat < 0 || c.CodCat > 999999 {
		return false, "Código de categoria inválido."
	}

	if c.CodSubcat.Valid {
		if c.CodSubcat.Int64 < 0 || c.CodSubcat.Int64 > 9999999 {
			return false, "Código de subcategoria inválido."
		}
	}

	if c.DiamInterno.Valid {
		if len(c.DiamInterno.String) > 10 {
			return false, "Diâmetro interno deve conter até 10 caracteres."
		}
	}

	if c.DiamExterno.Valid {
		if c.DiamExterno.Float64 > 99999.99 {
			return false, "Diâmetro externo deve ser inferior a 100.000."
		}

		c.DiamExterno.Float64 = float64(int(c.DiamExterno.Float64*100) / 100)
	}

	if c.DiamNominal.Valid {
		if len(c.DiamNominal.String) > 6 {
			return false, "Diâmetro nominal deve conter até 6 caracteres."
		}
	}

	if c.MedidaD.Valid {
		if c.MedidaD.Int64 > 999 {
			return false, "Medida D deve conter até 3 numerais."
		}
	}

	if c.PrensadoReusavel.Valid {
		if len(c.PrensadoReusavel.String) != 1 {
			return false, "Prensado reusável deve conter apenas 1 caractere."
		}
	}

	if c.Mangueira.Valid {
		if len(c.Mangueira.String) > 30 {
			return false, "Mangueira deve conter até 30 caracteres."
		}
	}

	if c.Material.Valid {
		if len(c.Material.String) > 20 {
			return false, "Material de componente deve conter até 20 caracteres."
		}
	}

	if c.Norma.Valid {
		if len(c.Norma.String) > 20 {
			return false, "Norma de componente deve conter até 20 caracteres."
		}
	}

	if c.Bitola.Valid {
		if c.Bitola.Int64 < -999 || c.Bitola.Int64 > 999 {
			return false, "Bitola deve conter até 3 caracteres."
		}
	}

	if c.ValorEntrada <= 0 || c.ValorEntrada > 999999999.99 {
		return false, "Valor de entrada deve ser positivo e inferior a 100 milhões."
	}

	if c.ValorSaida <= 0 || c.ValorSaida > 999999999.99 {
		return false, "Valor de saída deve ser positivo e inferior a 100 milhões."
	}

	c.ValorEntrada = float64(int(c.ValorEntrada*100) / 100)
	c.ValorSaida = float64(int(c.ValorSaida*100) / 100)

	return true, ""
}
