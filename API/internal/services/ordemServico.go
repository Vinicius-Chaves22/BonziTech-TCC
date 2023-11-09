// ordemServico.go possui todas as funcionalidades relacionadas
// às ordens de serviço do aplicativo da Connect
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodasOrdensServico(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM ordem_servico;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var ordensServico []models.OrdemServico
	for rows.Next() {
		var os models.OrdemServico
		err := rows.Scan(&os.CodOS, &os.DataEmissao, &os.CodCli, &os.NomeCli,
			&os.Pedido, &os.Concluida)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao retornar ordens de serviço."})
			return
		}

		ordensServico = append(ordensServico, os)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"ordensServico": ordensServico,
		"message":       "Ordens de serviço encontradas com sucesso!",
	})
}

func CriarOrdemServico(c *gin.Context) {
	var os models.OrdemServico
	err := c.BindJSON(&os)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	valido, erroOS := os.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroOS})
		return
	}

	os.NomeCli, err = getNomeCliente(os.CodCli)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado para criar ordem de serviço."})
		return
	}

	insert := `
		INSERT INTO ordem_servico (data_emissao, cod_cli, nome_cli, pedido, concluida)
		VALUES(?, ?, ?, ?, ?);`
	_, err = DB.Exec(insert, os.DataEmissao, os.CodCli, os.NomeCli,
		os.Pedido, os.Concluida)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar ordem de serviço."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Ordem de serviço criada com sucesso!"})
}

func AtualizarOrdemServico(c *gin.Context) {
	var os models.OrdemServico
	err := c.BindJSON(&os)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	valido, erroOS := os.EValida()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroOS})
		return
	}

	if !ordemServicoExiste(strconv.Itoa(os.CodOS)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Ordem de serviço não encontrada."})
		return
	}

	os.NomeCli, err = getNomeCliente(os.CodCli)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado para atualizar ordem de serviço."})
		return
	}

	update := `
		UPDATE ordem_servico SET data_emissao = ?, cod_cli = ?, nome_cli = ?,
		pedido = ?, concluida = ? WHERE cod_os = ?;`
	_, err = DB.Exec(update, os.DataEmissao, os.CodCli, os.NomeCli,
		os.Pedido, os.Concluida, os.CodOS)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar ordem de serviço."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ordem de serviço atualizada com sucesso!"})
}

func DeletarOrdemServico(c *gin.Context) {
	codOS := c.Param("codOS")

	if !ordemServicoExiste(codOS) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Ordem de serviço a ser excluída não encontrada."})
		return
	}

	_, err := DB.Exec("DELETE FROM ordem_servico WHERE cod_os = ?;", codOS)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir ordem de serviço."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ordem de serviço excluída com sucesso!"})
}

// Verifica se ordem de serviço existe
func ordemServicoExiste(codOS string) bool {
	row := DB.QueryRow("SELECT cod_os FROM ordem_servico WHERE cod_os = ?;", codOS)

	var codOSRow int
	if err := row.Scan(&codOSRow); err != nil {
		return false
	}

	return true
}

// Retorna o nome de um cliente dado seu código
func getNomeCliente(codCli int) (string, error) {
	row := DB.QueryRow("SELECT nome_cli FROM clientes WHERE cod_cli = ?;", codCli)

	var nomeCli string
	if err := row.Scan(&nomeCli); err != nil {
		return "", err
	}

	return nomeCli, nil
}
