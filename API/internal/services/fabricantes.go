// fabricantes.go possui todas as funcionalidades relacionadas
// a fabricantes do aplicativo da Connect
package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodosFabricantes(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM fabricantes;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados. Tente novamente mais tarde."})
		return
	}

	var fabs []models.Fabricante
	for rows.Next() {
		var fab models.Fabricante
		err := rows.Scan(&fab.CodFab, &fab.Nome, &fab.RazaoSocial, &fab.Telefone,
			&fab.Fax, &fab.Celular, &fab.NomeContato, &fab.Endereco, &fab.Cidade,
			&fab.Estado, &fab.CEP)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		fabs = append(fabs, fab)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"fabricantes": fabs,
		"message":     "Fabricantes encontrados com sucesso!",
	})
}

func MostrarFabricante(c *gin.Context) {
	codFab := c.Param("codFab")

	var fab models.Fabricante
	row := DB.QueryRow("SELECT * FROM fabricantes WHERE cod_fab = ?;", codFab)
	err := row.Scan(&fab.CodFab, &fab.Nome, &fab.RazaoSocial, &fab.Telefone,
		&fab.Fax, &fab.Celular, &fab.NomeContato, &fab.Endereco, &fab.Cidade,
		&fab.Estado, &fab.CEP)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Fabricante não encontrado."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"fabricante": fab,
		"message":    "Fabricante encontrado com sucesso!",
	})
}

func AdicionarFabricante(c *gin.Context) {
	var fab models.Fabricante
	if err := c.BindJSON(&fab); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de fabricante inválidos."})
		return
	}

	valido, erroFab := fab.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroFab})
		return
	}

	insert := `INSERT INTO fabricantes
		(nome_fab, razao_social, telefone, fax, celular, nome_contato, endereco, cidade, estado, cep)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := DB.Exec(insert, fab.Nome, fab.RazaoSocial, fab.Telefone, fab.Fax,
		fab.Celular, fab.NomeContato, fab.Endereco, fab.Cidade, fab.Estado, fab.CEP)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar fabricante. Verifique se todos os dados estão corretos e tente novamente."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Fabricante registrado com sucesso!"})
}

func AtualizarFabricante(c *gin.Context) {
	var fab models.Fabricante
	if err := c.BindJSON(&fab); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de fabricante inválidos."})
		return
	}

	valido, erroFab := fab.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroFab})
		return
	}

	// Verificar se fabricante a ser alterado existe
	rows := DB.QueryRow("SELECT cod_fab FROM fabricantes WHERE cod_fab = ?;", fab.CodFab)

	var codFabRows int
	err := rows.Scan(&codFabRows)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Fabricante não existe, ou não pode ser atualizado."})
		return
	}

	update := `
		UPDATE fabricantes SET nome_fab = ?, nome_contato = ?, razao_social = ?,
		telefone = ?, celular = ?, fax = ?, endereco = ?, cidade = ?, estado = ?,
		cep = ? WHERE cod_fab = ?;`
	_, err = DB.Exec(update, fab.Nome, fab.NomeContato, fab.RazaoSocial,
		fab.Telefone, fab.Celular, fab.Fax, fab.Endereco, fab.Cidade, fab.Estado,
		fab.CEP, fab.CodFab)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao atualizar fabricante. Verifique se todos os dados estão corretos e tente novamente."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Fabricante atualizado com sucesso!"})
}

func DeletarFabricante(c *gin.Context) {
	codFab := c.Param("codFab")

	// Verificar se fabricante a ser excluído existe
	rows := DB.QueryRow("SELECT cod_fab FROM fabricantes WHERE cod_fab = ?;", codFab)

	var codFabRows int
	err := rows.Scan(&codFabRows)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Fabricante não existe, ou não pode ser excluído."})
		return
	}

	delete := "DELETE FROM fabricantes WHERE cod_fab = ?;"
	_, err = DB.Exec(delete, codFab)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir fabricante. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Fabricante excluído com sucesso!"})
}
