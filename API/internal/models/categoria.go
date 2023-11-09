package models

type Categoria struct {
	CodCat     int    `json:"codCat"`
	NomeCat    string `json:"nomeCat"`
	UnidMedida string `json:"unidMedida"`
	Montagem   bool   `json:"montagem"`
	Apelido    string `json:"apelido"`
}

func (c Categoria) EValida() (bool, string) {
	if c.CodCat < 0 || c.CodCat > 999999 {
		return false, "Código de categoria inválido."
	}

	if c.NomeCat == "" || len(c.NomeCat) > 30 {
		return false, "Nome de categoria deve conter de 1 a 30 caracteres."
	}

	if c.UnidMedida == "" || len(c.UnidMedida) > 3 {
		return false, "Unidade de medida deve conter de 1 a 3 caracteres. Ex: 'cm'"
	}

	if c.Apelido == "" || len(c.Apelido) > 4 {
		return false, "Apelido de categoria deve conter de 1 a 4 caracteres."
	}

	return true, ""
}
