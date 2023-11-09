package models

type Subcategoria struct {
	CodSubcat int    `json:"codSubcat"`
	CodCat    int    `json:"codCat"`
	Nome      string `json:"nome"`
}

func (s Subcategoria) EValida() (bool, string) {
	if s.CodSubcat < 0 || s.CodSubcat > 9999999 {
		return false, "Código de subcategoria inválido."
	}

	if s.CodCat < 0 || s.CodCat > 999999 {
		return false, "Código de categoria inválido."
	}

	if s.Nome == "" || len(s.Nome) > 60 {
		return false, "Nome de subcategoria deve conter de 1 a 60 caracteres."
	}

	return true, ""
}
