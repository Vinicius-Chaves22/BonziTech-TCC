// clientes.go possui todas as funcionalidades relacionadas
// a clientes do aplicativo da Connect
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func MostrarTodosClientes(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM clientes;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var clientes []models.Cliente
	for rows.Next() {
		var cli models.Cliente
		err := rows.Scan(&cli.CodCli, &cli.NomeEmpresa, &cli.NomeCli, &cli.Tipo,
			&cli.DiaReg, &cli.Endereco, &cli.Bairro, &cli.Cidade, &cli.Estado,
			&cli.CEP, &cli.Email)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao retornar clientes"})
			return
		}

		clientes = append(clientes, cli)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"clientes": clientes,
		"message":  "Clientes encontrados com sucesso!",
	})
}

func MostrarCliente(c *gin.Context) {
	codCli := c.Param("codCli")

	selectCli := "SELECT * FROM clientes WHERE cod_cli = ?;"

	var cli models.Cliente
	row := DB.QueryRow(selectCli, codCli)
	err := row.Scan(&cli.CodCli, &cli.NomeEmpresa, &cli.NomeCli, &cli.Tipo,
		&cli.DiaReg, &cli.Endereco, &cli.Bairro, &cli.Cidade, &cli.Estado,
		&cli.CEP, &cli.Email)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}

	telefones, err := getTelefonesCliente(codCli)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao retornar telefones do cliente."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"cliente":   cli,
		"telefones": telefones,
		"message":   "Cliente encontrado com sucesso!",
	})
}

func CriarCliente(c *gin.Context) {
	var cli models.Cliente
	if err := c.BindJSON(&cli); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de cliente inválidos"})
		return
	}

	valido, erroCli := cli.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroCli})
		return
	}

	insert := `
		INSERT INTO clientes (nome_empresa, nome_cli, tipo, dia_reg, endereco, 
		bairro, cidade, estado, cep, email) VALUES(?, ?, ?, CURRENT_DATE(), ?, 
		?, ?, ?, ?, ?);`

	_, err := DB.Exec(insert, cli.NomeEmpresa, cli.NomeCli, cli.Tipo, cli.Endereco,
		cli.Bairro, cli.Cidade, cli.Estado, cli.CEP, cli.Email)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cliente criado com sucesso!"})
}

func AtualizarCliente(c *gin.Context) {
	var cli models.Cliente
	if err := c.BindJSON(&cli); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de cliente inválidos"})
		return
	}

	valido, erroCli := cli.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroCli})
		return
	}

	if !clienteExiste(strconv.Itoa(cli.CodCli)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente a ser atualizado não encontrado."})
		return
	}

	update := `
		UPDATE clientes SET nome_empresa = ?, nome_cli = ?, tipo = ?, endereco = ?, 
		bairro = ?, cidade = ?, estado = ?, cep = ?, email = ? WHERE cod_cli = ?;`

	_, err := DB.Exec(update, cli.NomeEmpresa, cli.NomeCli, cli.Tipo, cli.Endereco,
		cli.Bairro, cli.Cidade, cli.Estado, cli.CEP, cli.Email, cli.CodCli)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cliente criado com sucesso!"})
}

func DeletarCliente(c *gin.Context) {
	codCli := c.Param("codCli")

	if !clienteExiste(codCli) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente a ser excluído não encontrado."})
		return
	}

	_, err := DB.Exec("DELETE FROM clientes WHERE cod_cli = ?;", codCli)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	_, err = DB.Exec("DELETE FROM telefones WHERE cod_cli = ?;", codCli)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Cliente excluído com sucesso!"})
}

func CadastrarTelefone(c *gin.Context) {
	var telefones []models.Telefone
	if err := c.BindJSON(&telefones); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	insert := `
		INSERT INTO telefones (cod_cli, telefone, nome_tel, tipo_contato, tipo_cli)
		VALUES (?, ?, ?, ?, ?);`

	for _, tel := range telefones {
		valido, erroTel := tel.EValido()
		if !valido {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroTel})
			return
		}

		if !clienteExiste(strconv.Itoa(tel.CodCli)) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente em posse de telefone não encontrado."})
			return
		}

		_, err := DB.Exec(insert, tel.CodCli, tel.Telefone, tel.NomeTel,
			tel.TipoContato, tel.TipoCli)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar telefones."})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Telefones registrados com sucesso!"})
}

func AtualizarTelefone(c *gin.Context) {
	var tel models.Telefone
	if err := c.BindJSON(&tel); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos."})
		return
	}

	valido, erroTel := tel.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroTel})
		return
	}

	if !clienteExiste(strconv.Itoa(tel.CodCli)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cliente em posse de telefone não encontrado."})
		return
	}

	if !telefoneExiste(strconv.Itoa(tel.CodTel)) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Telefone a ser atualizado não encontrado."})
		return
	}

	update := `
		UPDATE telefones SET cod_cli = ?, telefone = ?, nome_tel = ?,
		tipo_contato = ?, tipo_cli = ? WHERE cod_tel = ?;`
	_, err := DB.Exec(update, tel.CodCli, tel.Telefone, tel.NomeTel,
		tel.TipoContato, tel.TipoCli)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar telefone."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Telefone atualizado com sucesso!"})
}

func DeletarTelefone(c *gin.Context) {
	codTel := c.Param("codTel")

	if !telefoneExiste(codTel) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Telefone a ser excluído não encontrado."})
		return
	}

	_, err := DB.Exec("DELETE FROM telefones WHERE cod_tel = ?;", codTel)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir telefone de cliente."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Telefone excluído com sucesso!"})
}

// Verifica se cliente existe
func clienteExiste(codCli string) bool {
	row := DB.QueryRow("SELECT cod_cli FROM clientes WHERE cod_cli = ?;", codCli)

	var codCliRow int
	if err := row.Scan(&codCliRow); err != nil {
		return false
	}

	return true
}

// Verifica se telefone está cadastrado
func telefoneExiste(codTel string) bool {
	row := DB.QueryRow("SELECT cod_cli FROM telefones WHERE cod_tel = ?;", codTel)

	var codTelRow int
	if err := row.Scan(&codTelRow); err != nil {
		return false
	}

	return true
}

// Retorna todos os telefones de um cliente
func getTelefonesCliente(codCli string) ([]models.Telefone, error) {
	rows, err := DB.Query("SELECT * FROM telefones WHERE cod_cli = ?;", codCli)
	if err != nil {
		return nil, err
	}

	var telefones []models.Telefone
	for rows.Next() {
		var tel models.Telefone
		err := rows.Scan(&tel.CodTel, &tel.CodCli, &tel.Telefone, &tel.NomeTel,
			&tel.TipoContato, &tel.TipoCli)
		if err != nil {
			return telefones, err
		}

		telefones = append(telefones, tel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	return telefones, nil
}
