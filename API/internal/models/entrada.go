package models

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Entrada struct {
	CodEntd    int         `json:"codEntd"`
	CodFab     null.Int    `json:"codFab"`
	NomeFab    string      `json:"nomeFab"`
	DataVenda  string      `json:"dataVenda"`
	NotaFiscal null.String `json:"notaFiscal"`
	ValorTotal float64     `json:"valorTotal"`
}

func (e Entrada) EValida() (bool, string) {
	if e.CodEntd < 0 || e.CodEntd > 99999999 {
		return false, "Código de entrada inválido."
	}

	if !e.CodFab.Valid {
		return false, "Código de fabricante não pode ser nulo."
	}

	if e.CodFab.Int64 <= 0 || e.CodFab.Int64 > 999999 {
		return false, "Código de fabricante deve estar entre 1 e 1.000.000."
	}

	if e.DataVenda == "" {
		e.DataVenda = time.Now().Format("2006-01-02")
	}

	if e.ValorTotal <= 0 || e.ValorTotal > 999999999.99 {
		return false, "Valor total de entrada deve ser positivo e menor que 1 bilhão."
	}

	e.ValorTotal = float64(int(e.ValorTotal*100) / 100)

	return true, ""
}
