// categorias.go possui todas as funcionalidades relacionadas
// às categorias de componentes
package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodasCategorias(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM categorias;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar ao banco de dados. Tente novamente mais tarde."})
		return
	}

	var cats []models.Categoria
	for rows.Next() {
		var cat models.Categoria
		err := rows.Scan(&cat.CodCat, &cat.NomeCat, &cat.UnidMedida,
			&cat.Montagem, &cat.Apelido)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		cats = append(cats, cat)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"categorias": cats,
		"message":    "Categorias encontradas com sucesso!",
	})
}

func MostrarComponentesCategoria(c *gin.Context) {
	codCat := c.Param("codCat")

	var cat models.Categoria
	row := DB.QueryRow("SELECT * FROM categorias WHERE cod_cat = ?;", codCat)
	err := row.Scan(&cat.CodCat, &cat.NomeCat, &cat.UnidMedida, &cat.Montagem,
		&cat.Apelido)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Erro ao encontrar categoria."})
		return
	}

	rows, err := DB.Query("SELECT * FROM componentes WHERE cod_cat = ?;", codCat)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados. Tente novamente mais tarde."})
		return
	}

	var componentes []models.Componente
	for rows.Next() {
		var comp models.Componente
		err := rows.Scan(
			&comp.CodComp, &comp.CodPeca, &comp.Especificacao, &comp.CodCat,
			&comp.CodSubcat, &comp.DiamInterno, &comp.DiamExterno,
			&comp.DiamNominal, &comp.MedidaD, &comp.Costura,
			&comp.PrensadoReusavel, &comp.Mangueira, &comp.Material,
			&comp.Norma, &comp.Bitola, &comp.ValorEntrada, &comp.ValorSaida)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados. Tente novamente mais tarde."})
			return
		}

		componentes = append(componentes, comp)
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"componentes": componentes,
		"categoria":   cat,
		"message":     "Componentes de categoria encontrados com sucesso!",
	})
}

func CriarCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := c.BindJSON(&cat); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de categoria inválidos."})
		return
	}

	catEValida, erroCat := cat.EValida()
	if !catEValida {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroCat})
		return
	}

	insert := "INSERT INTO categorias (nome_cat, unid_medida, montagem, apelido) VALUES(?, ?, ?, ?);"
	_, err := DB.Exec(insert, cat.NomeCat, cat.UnidMedida, cat.Montagem, cat.Apelido)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar categoria. Verifique se uma categoria com o mesmo nome já não existe e tente novamente."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Categoria criada com sucesso!"})
}

func AtualizarCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := c.BindJSON(&cat); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de categoria inválidos."})
		return
	}

	catEValida, erroCat := cat.EValida()
	if !catEValida {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroCat})
		return
	}

	row := DB.QueryRow("SELECT nome_cat FROM categorias WHERE cod_cat = ?;", cat.CodCat)

	var nomeCat string
	if err := row.Scan(&nomeCat); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Categoria não existe."})
		return
	}

	update := "UPDATE categorias SET nome_cat = ?, unid_medida = ?, montagem = ?, apelido = ?;"
	_, err := DB.Exec(update, cat.NomeCat, cat.UnidMedida, cat.Montagem, cat.Apelido)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar categoria. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Categoria atualizada com sucesso!"})
}

func DeletarCategoria(c *gin.Context) {
	codCat := c.Param("codCat")

	row := DB.QueryRow("SELECT nome_cat FROM categorias WHERE cod_cat = ?;", codCat)

	var nomeCat string
	if err := row.Scan(&nomeCat); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Categoria não existe."})
		return
	}

	delete := "DELETE FROM categorias WHERE cod_cat = ?;"
	_, err := DB.Exec(delete, codCat)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir categoria. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Categoria excluída com sucesso!"})
}
