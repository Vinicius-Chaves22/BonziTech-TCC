// vendas.go possui todas as funcionalidades relacionadas
// a vendas do aplicativo da Connect
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodasVendas(c *gin.Context) {
	selectVenda := `
		SELECT vendas.*, clientes.nome_cli
		FROM vendas
		INNER JOIN clientes
		ON vendas.cod_cli = clientes.cod_cli;`

	rows, err := DB.Query(selectVenda) // TODO: mudar para inner join
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var vendas []models.Venda
	for rows.Next() {
		var v models.Venda
		err := rows.Scan(&v.CodVenda, &v.DataVenda, &v.CodCli, &v.NomeCli, &v.CodOS,
			&v.ValorTotal, &v.Descricao, &v.NomeCli)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		vendas = append(vendas, v)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"vendas":  vendas,
		"message": "Vendas encontradas com sucesso!",
	})
}

func MostrarVenda(c *gin.Context) {
	codVenda := c.Param("codVenda")

	row := DB.QueryRow(`
		SELECT vendas.*, clientes.nome_cli
		FROM vendas
		INNER JOIN clientes
		ON vendas.cod_cli = clientes.cod_cli
		WHERE cod_venda = ?;`,
		codVenda)

	var v models.Venda
	err := row.Scan(&v.CodVenda, &v.DataVenda, &v.CodCli, &v.NomeCli, &v.CodOS,
		&v.ValorTotal, &v.Descricao, &v.NomeCli)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Venda não encontrada."})
		return
	}

	componentesSaida, err := getComponentesSaida(codVenda)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Não é possível receber componentes de venda."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"venda":       v,
		"componentes": componentesSaida,
		"message":     "Venda encontrada com sucesso!",
	})
}

func AdicionarVenda(c *gin.Context) {
	var venda models.Venda
	if err := c.BindJSON(&venda); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de venda inválidos."})
		return
	}

	valido, erroVenda := venda.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroVenda})
		return
	}

	insert := `
		INSERT INTO vendas (data_venda, cod_cli, cod_os, valor_total, descricao)
		VALUES(?, ?, ?, 0, ?);`
	_, err := DB.Exec(insert, venda.DataVenda, venda.CodCli, venda.CodOS,
		venda.Descricao)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar nova venda."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Venda cadastrada com sucesso!"})
}

func AdicionarComponentesVenda(c *gin.Context) {

	// TODO: remover itens vendidos de estoque
}

func AtualizarVenda(c *gin.Context) {
	var venda models.Venda
	if err := c.BindJSON(&venda); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de venda inválidos."})
		return
	}

	valido, erroVenda := venda.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroVenda})
		return
	}

	if !vendaExiste(strconv.Itoa(venda.CodVenda)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Venda a ser editada não encontrada."})
		return
	}

	update := `
		UPDATE vendas SET data_venda = ?, cod_cli = ?, cod_os = ?,
		descricao = ? WHERE cod_venda = ?;`
	_, err := DB.Exec(update, venda.DataVenda, venda.CodCli, venda.CodOS,
		venda.Descricao, venda.CodVenda)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar venda."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Venda editada com sucesso!"})
}

func AtualizarComponenteVenda(c *gin.Context) {
	// TODO
}

func DeletarVenda(c *gin.Context) {
	codVenda := c.Param("codVenda")

	if !vendaExiste(codVenda) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Venda a ser excluída não encontrada."})
		return
	}

	delete := "DELETE FROM vendas WHERE cod_venda = ?;"
	_, err := DB.Exec(delete, codVenda)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir venda."})
		return
	}

	delete = "DELETE FROM componentes_saida WHERE cod_venda = ?;"
	_, err = DB.Exec(delete, codVenda)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir componentes de venda."})
		return
	}

	// TODO: retornar itens ao estoque

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Venda excluída com sucesso!"})
}

// Retorna todos os componentes de uma dada venda
func getComponentesSaida(codVenda string) ([]models.ComponenteSaida, error) {
	rows, err := DB.Query("SELECT * FROM componentes_saida WHERE cod_venda = ?;", codVenda)
	if err != nil {
		return nil, err
	}

	var componentesSaida []models.ComponenteSaida
	for rows.Next() {
		var compSaida models.ComponenteSaida
		err := rows.Scan(&compSaida.CodCompVenda, &compSaida.CodVenda,
			&compSaida.CodComp, &compSaida.ValorUnit)
		if err != nil {
			return nil, err
		}

		componentesSaida = append(componentesSaida, compSaida)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	return componentesSaida, nil
}

// Verifica se venda existe
func vendaExiste(codVenda string) bool {
	row := DB.QueryRow("SELECT cod_venda FROM vendas WHERE cod_venda = ?;", codVenda)

	var codVendaRow int
	if err := row.Scan(&codVendaRow); err != nil {
		return false
	}

	return true
}
