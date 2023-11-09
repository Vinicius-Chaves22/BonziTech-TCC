// estoque.go possui todas as funcionalidades relacionadas
// a estoque do aplicativo da Connect
package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarEstoque(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM estoque;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar ao banco de dados. Tente novamente mais tarde."})
		return
	}

	var estoque []models.Estoque
	for rows.Next() {
		var e models.Estoque
		err := rows.Scan(&e.CodEstq, &e.CodComp, &e.QuantMin, &e.QuantMax,
			&e.QuantAtual)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao retornar estoque. Tente novamente mais tarde."})
			return
		}

		estoque = append(estoque, e)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"estoque": estoque,
		"message": "Estoque encontrado com sucesso!",
	})
}

func AdicionarComponenteEstoque(c *gin.Context) {
	var e models.Estoque
	if err := c.BindJSON(&e); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de estoque inválidos."})
		return
	}

	valido, erroEstq := e.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroEstq})
		return
	}

	// TODO: validar e.CodComp

	insert := "INSERT INTO estoque (cod_comp, quant_min, quant_max, quantidade) VALUES(?, ?, ?, ?);"
	_, err := DB.Exec(insert, e.CodComp, e.QuantMin, e.QuantMax, e.QuantAtual)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao adicionar componente ao estoque. Cheque se o componente já foi adicionado ou se ele não existe e tente novamente."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Componente adicionado ao estoque com sucesso!"})
}

func AtualizarEstoque(c *gin.Context) {
	var e models.Estoque
	if err := c.BindJSON(&e); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de estoque inválidos."})
		return
	}

	valido, erroEstq := e.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroEstq})
		return
	}

	row := DB.QueryRow("SELECT cod_estq FROM estoque WHERE cod_comp = ?;", e.CodComp)

	var codEstqRow int
	if err := row.Scan(&codEstqRow); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Componente não existe no estoque."})
		return
	}

	update := "UPDATE estoque SET quant_min = ?, quant_max = ?, quantidade = ? WHERE cod_comp = ?;"
	_, err := DB.Exec(update, e.QuantMin, e.QuantMax, e.QuantAtual, e.CodComp)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar dados no estoque. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Dados de componente atualizado no estoque com sucesso!"})
}

func DeletarComponenteEstoque(c *gin.Context) {
	codComp := c.Param("codComp")

	row := DB.QueryRow("SELECT cod_estq FROM estoque WHERE cod_comp = ?;", codComp)

	var codEstqRow int
	if err := row.Scan(&codEstqRow); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Componente não existe no estoque."})
		return
	}

	delete := "DELETE FROM estoque WHERE cod_comp = ?;"
	_, err := DB.Exec(delete, codComp)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir dados do estoque. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Dados de componente excluídos do estoque com sucesso!"})
}
