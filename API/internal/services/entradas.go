// entradas.go possui todas as funcionalidades relacionadas
// à entrada de componentes na empresa Connect
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodasEntradas(c *gin.Context) {
	selectEntd := `
		SELECT entradas.*, fabricantes.nome_fab
		FROM entradas
		INNER JOIN fabricantes ON entradas.cod_fab = fabricantes.cod_fab;
		`

	rows, err := DB.Query(selectEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var entradas []models.Entrada
	for rows.Next() {
		var entd models.Entrada
		err := rows.Scan(&entd.CodEntd, &entd.CodFab, &entd.DataVenda,
			&entd.NotaFiscal, &entd.ValorTotal, &entd.NomeFab)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		entradas = append(entradas, entd)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"entradas": entradas,
		"message":  "Entradas encontradas com sucesso!",
	})
}

func MostrarEntrada(c *gin.Context) {
	codEntd := c.Param("codEntd")

	row := DB.QueryRow(`
		SELECT entradas.*, fabricantes.nome_fab
		FROM entradas
		INNER JOIN fabricantes ON entradas.cod_fab = fabricantes.cod_fab
		WHERE cod_entd = ?;`,
		codEntd)

	var entd models.Entrada
	err := row.Scan(&entd.CodEntd, &entd.CodFab, &entd.DataVenda,
		&entd.NotaFiscal, &entd.ValorTotal, &entd.NomeFab)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entrada não encontrada."})
		return
	}

	componentesEntd, err := getComponentesEntrada(codEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"entrada":     entd,
		"componentes": componentesEntd,
		"message":     "Entrada encontrada com sucesso!",
	})
}

func AdicionarEntrada(c *gin.Context) {
	var entd models.Entrada
	if err := c.BindJSON(&entd); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos."})
		return
	}

	valido, erroEntd := entd.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroEntd})
		return
	}

	if !fabricanteExiste(int(entd.CodFab.Int64)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Fornecedor de componente não encontrado no sistema. Cadastre-o e tente novamente."})
		return
	}

	insert := `
		INSERT INTO entradas (cod_fab, data_venda, nota_fiscal, valor_total)
		VALUES(?, ?, ?, ?);`

	_, err := DB.Exec(insert, entd.CodFab, entd.DataVenda,
		entd.NotaFiscal, entd.ValorTotal)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar entrada."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Entrada adicionada com sucesso!"})
}

func AdicionarComponentesEntrada(c *gin.Context) {
	var componentesEntd []models.ComponenteEntrada
	if err := c.BindJSON(&componentesEntd); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	codEntd := componentesEntd[0].CodEntd

	if !entradaExiste(strconv.Itoa(codEntd)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entrada não encontrada."})
		return
	}

	insert := `
		INSERT INTO componentes_entrada (cod_entd, cod_comp, quantidade,
		valor_unit) VALUES(?, ?, ?, ?);`

	for _, comp := range componentesEntd {
		if codEntd != comp.CodEntd {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Você só pode atualizar os componentes de uma entrada por vez."})
			return
		}

		valido, erroComp := comp.EValido()
		if !valido {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroComp})
			return
		}

		_, err := DB.Exec(insert, comp.CodEntd, comp.CodComp, comp.Quantidade,
			comp.ValorUnit)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar componentes."})
			return
		}

		// TODO: Inserir em estoque se não existir ainda

		// Atualizar estoque
		update := "UPDATE estoque SET quantidade = quantidade + ? WHERE cod_comp = ?;"
		_, err = DB.Exec(update, comp.Quantidade, comp.CodComp)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar itens no estoque."})
			return
		}
	}

	err := atualizarValorTotalEntrada(codEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular valor total de entrada."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Componentes criados com sucesso!"})
}

func AtualizarEntrada(c *gin.Context) {
	var entd models.Entrada
	if err := c.BindJSON(&entd); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos."})
		return
	}

	valido, erroEntd := entd.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroEntd})
		return
	}

	if !entradaExiste(strconv.Itoa(entd.CodEntd)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entrada a ser atualizada não encontrada."})
		return
	}

	if !fabricanteExiste(int(entd.CodFab.Int64)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Fornecedor de componente não encontrado no sistema. Cadastre-o e tente novamente."})
		return
	}

	update := `
		UPDATE entradas SET cod_fab = ?, data_venda = ?, nota_fiscal = ?
		WHERE cod_entd = ?;`
	_, err := DB.Exec(update, entd.CodFab, entd.DataVenda, entd.NotaFiscal,
		entd.CodEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar entrada. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Entrada atualizada com sucesso!"})
}

func AtualizarComponenteEntrada(c *gin.Context) {
	var compEntd models.ComponenteEntrada
	if err := c.BindJSON(&compEntd); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de componente inválidos."})
		return
	}

	compEValido, erroCompEntd := compEntd.EValido()
	if !compEValido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroCompEntd})
		return
	}

	row := DB.QueryRow(
		"SELECT cod_entd FROM componentes_entrada WHERE cod_comp_entd = ?;",
		compEntd.CodCompEntd)

	var codEntdRow int
	if err := row.Scan(&codEntdRow); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Componente a ser alterado não encontrado."})
		return
	}

	if codEntdRow != compEntd.CodEntd {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Entrada do componente não pode ser alterada."})
		return
	}

	update := `
		UPDATE componentes_entrada SET cod_comp = ?, quantidade = ?,
		valor_unit = ? WHERE cod_comp_entd = ?;`
	_, err := DB.Exec(update, compEntd.CodComp, compEntd.Quantidade,
		compEntd.ValorUnit, compEntd.CodCompEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar componente. Tente novamente mais tarde."})
		return
	}

	// TODO: Atualizar estoque

	err = atualizarValorTotalEntrada(compEntd.CodEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular valor total de entrada."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Componente atualizado com sucesso!"})
}

func DeletarEntrada(c *gin.Context) {
	codEntd := c.Param("codEntd")

	if !entradaExiste(codEntd) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entrada a ser exluída não encontrada."})
		return
	}

	delete := "DELETE FROM entradas WHERE cod_entd = ?;"
	_, err := DB.Exec(delete, codEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir entrada."})
		return
	}

	delete = "DELETE FROM componentes_entrada WHERE cod_entd = ?;"
	_, err = DB.Exec(delete, codEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir componentes da entrada."})
		return
	}

	// TODO: remover componentes de estoque

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Entrada excluída com sucesso!"})
}

func DeletarComponenteEntrada(c *gin.Context) {
	codCompEntd := c.Param("codCompEntd")

	selectComp := "SELECT cod_entd FROM componentes_entrada WHERE cod_comp_entd = ?;"
	row := DB.QueryRow(selectComp, codCompEntd)

	var codEntd int
	if err := row.Scan(&codEntd); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Componente a ser excluído não encontrado."})
		return
	}

	delete := "DELETE FROM componentes_entrada WHERE cod_comp_entd = ?;"
	_, err := DB.Exec(delete, codCompEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir componente de entrada."})
		return
	}

	err = atualizarValorTotalEntrada(codEntd)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular valor total de entrada."})
		return
	}

	// TODO: remover componentes de estoque

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Componente excluído com sucesso!"})
}

// Função que retorna componentes de uma dada entrada por meio de seu código
func getComponentesEntrada(codEntd string) ([]models.ComponenteEntrada, error) {
	rows, err := DB.Query("SELECT * FROM componentes_entrada WHERE cod_entd = ?;", codEntd)
	if err != nil {
		return nil, err
	}

	var componentesEntd []models.ComponenteEntrada
	for rows.Next() {
		var compEntd models.ComponenteEntrada
		err := rows.Scan(&compEntd.CodCompEntd, &compEntd.CodEntd,
			&compEntd.CodComp, &compEntd.Quantidade, &compEntd.ValorUnit)
		if err != nil {
			return nil, err
		}

		componentesEntd = append(componentesEntd, compEntd)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	return componentesEntd, nil
}

// Calcula o valor total pago em uma entrada a partir de seus componentes
func atualizarValorTotalEntrada(codEntd int) error {
	componentes, err := getComponentesEntrada(strconv.Itoa(codEntd))
	if err != nil {
		return err
	}

	var valorTotal float64
	for _, comp := range componentes {
		valorTotal += comp.Quantidade * comp.ValorUnit
	}

	_, err = DB.Exec("UPDATE entradas SET valor_total = ? WHERE cod_entd = ?;",
		valorTotal, codEntd)
	if err != nil {
		return err
	}

	return nil
}

// Verifica se entrada com dado código existe
func entradaExiste(codEntd string) bool {
	row := DB.QueryRow("SELECT cod_entd FROM entradas WHERE cod_entd = ?;", codEntd)

	var codEntdRow int
	if err := row.Scan(&codEntdRow); err != nil {
		return false
	}

	return true
}

// Verifica se fabricante com dado código existe
func fabricanteExiste(codFab int) bool {
	row := DB.QueryRow("SELECT cod_fab FROM fabricantes WHERE cod_fab = ?;", codFab)

	var codFabRow int
	if err := row.Scan(&codFabRow); err != nil {
		return false
	}

	return true
}
