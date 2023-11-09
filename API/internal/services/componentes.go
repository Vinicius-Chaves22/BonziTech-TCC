// componentes.go possui todas as funcionalidades relacionadas
// aos componentes do aplicativo da Connect
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodosComponentes(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM componentes;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var comps []models.Componente
	for rows.Next() {
		var cmp models.Componente
		err := rows.Scan(&cmp.CodComp, &cmp.CodPeca, &cmp.Especificacao,
			&cmp.CodCat, &cmp.CodSubcat, &cmp.DiamInterno, &cmp.DiamExterno,
			&cmp.DiamNominal, &cmp.MedidaD, &cmp.Costura, &cmp.PrensadoReusavel,
			&cmp.Mangueira, &cmp.Material, &cmp.Norma, &cmp.Bitola,
			&cmp.ValorEntrada, &cmp.ValorSaida)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		comps = append(comps, cmp)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"componentes": comps,
		"message":     "Componentes encontrados com sucesso!",
	})
}

func MostrarComponente(c *gin.Context) {
	codComp := c.Param("codComp")

	row := DB.QueryRow("SELECT * FROM componentes WHERE cod_comp = ?;", codComp)

	var cmp models.Componente
	err := row.Scan(&cmp.CodComp, &cmp.CodPeca, &cmp.Especificacao,
		&cmp.CodCat, &cmp.CodSubcat, &cmp.DiamInterno, &cmp.DiamExterno,
		&cmp.DiamNominal, &cmp.MedidaD, &cmp.Costura, &cmp.PrensadoReusavel,
		&cmp.Mangueira, &cmp.Material, &cmp.Norma, &cmp.Bitola,
		&cmp.ValorEntrada, &cmp.ValorSaida)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Componente não encontrado."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"componente": cmp,
		"message":    "Componente encontrado com sucesso!",
	})
}

func AdicionarComponente(c *gin.Context) {
	var cmp models.Componente
	if err := c.BindJSON(&cmp); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	valido, erroComp := cmp.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroComp})
		return
	}

	insert := `INSERT INTO componentes (cod_peca, especificacao, cod_cat, cod_subcat,
		diam_interno, diam_externo, diam_nominal, medida_d, costura, prensado_reusavel,
		mangueira, material, norma, bitola, valor_entrada, valor_saida) VALUES(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := DB.Exec(insert, cmp.CodPeca, cmp.Especificacao,
		cmp.CodCat, cmp.CodSubcat, cmp.DiamInterno, cmp.DiamExterno,
		cmp.DiamNominal, cmp.MedidaD, cmp.Costura, cmp.PrensadoReusavel,
		cmp.Mangueira, cmp.Material, cmp.Norma, cmp.Bitola,
		cmp.ValorEntrada, cmp.ValorSaida)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar componente. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Componente cadastrado com sucesso!"})
}

func AtualizarComponente(c *gin.Context) {
	var cmp models.Componente
	if err := c.BindJSON(&cmp); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	valido, erroComp := cmp.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroComp})
		return
	}

	if !componenteExiste(strconv.Itoa(cmp.CodComp)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Componente a ser atualizado não encontrado."})
		return
	}

	update := `UPDATE componentes SET cod_peca = ?, especificacao = ?, cod_cat = ?, 
		cod_subcat = ?, diam_interno = ?, diam_externo = ?, diam_nominal = ?, 
		medida_d = ?, costura = ?, prensado_reusavel = ?, mangueira = ?, 
		material = ?, norma = ?, bitola = ?, valor_entrada = ?, valor_saida = ?
		WHERE cod_comp = ?;`
	_, err := DB.Exec(update, cmp.CodPeca, cmp.Especificacao,
		cmp.CodCat, cmp.CodSubcat, cmp.DiamInterno, cmp.DiamExterno,
		cmp.DiamNominal, cmp.MedidaD, cmp.Costura, cmp.PrensadoReusavel,
		cmp.Mangueira, cmp.Material, cmp.Norma, cmp.Bitola,
		cmp.ValorEntrada, cmp.ValorSaida, cmp.CodComp)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar componente. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Componente atualizado com sucesso!"})
}

func DeletarComponente(c *gin.Context) {
	codComp := c.Param("codComp")

	if !componenteExiste(codComp) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Componente a ser excluído não encontrado."})
		return
	}

	delete := "DELETE FROM componentes WHERE cod_comp = ?;"
	_, err := DB.Exec(delete, codComp)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Erro ao excluir componente. Tente novamente mais tarde."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Componente exlucído com sucesso!"})
}

// Verifica se componente existe no banco de dados
func componenteExiste(codComp string) bool {
	row := DB.QueryRow("SELECT cod_peca FROM componentes WHERE cod_comp = ?;", codComp)

	var codPeca string
	if err := row.Scan(&codPeca); err != nil {
		return false
	}

	return true
}
