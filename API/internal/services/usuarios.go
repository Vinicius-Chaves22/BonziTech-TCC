// usuarios.go contém as funcionalidades de manejo
// de usuários da aplicação
package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
	utils "github.com/vidacalura/BonziTech-TCC/internal/utils"
)

func MostrarUsuario(c *gin.Context) {
	codUsu := c.Param("codUsu")

	rows := DB.QueryRow("SELECT cod_usu, permissoes, nome, email, ativo FROM usuarios WHERE cod_usu = ?;", codUsu)

	var u models.Usuario
	err := rows.Scan(&u.CodUsuario, &u.Permissoes, &u.Nome, &u.Email, &u.Ativo)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"usuario": u,
		"message": "Usuário encontrado com sucesso!",
	})
}

func MostrarTodosUsuarios(c *gin.Context) {
	rows, err := DB.Query("SELECT cod_usu, permissoes, nome, email, ativo FROM usuarios WHERE ativo = TRUE;")
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(&u.CodUsuario, &u.Permissoes, &u.Nome, &u.Email, &u.Ativo)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
			return
		}

		usuarios = append(usuarios, u)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados."})
		return
	}

	defer rows.Close()

	c.IndentedJSON(http.StatusOK, gin.H{
		"usuarios": usuarios,
		"message":  "Usuários encontrados com sucesso!",
	})
}

func ValidarDadosLogin(c *gin.Context) {
	var u models.Usuario
	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de usuário inválidos."})
		return
	}

	if u.Email == "" || u.Senha == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email e senha não podem estar vazios."})
		return
	}

	u.Senha = utils.CriptografarSenha(u.Senha)

	query := "SELECT cod_usu FROM usuarios WHERE BINARY email = ? AND BINARY senha = ? AND ativo = TRUE;"
	rows := DB.QueryRow(query, u.Email, u.Senha)

	var codUsu int
	err = rows.Scan(&codUsu)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuário ou senha incorretos."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"codUsuario": codUsu,
		"message":    "Usuário encontrado com sucesso!",
	})
}

func AdicionarUsuario(c *gin.Context) {
	var novoUsuario models.Usuario
	err := c.BindJSON(&novoUsuario)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Dados de usuário inválidos."})
		return
	}

	valido, erroUsu := novoUsuario.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroUsu})
		return
	}

	novoUsuario.Senha = utils.CriptografarSenha(novoUsuario.Senha)

	insert := "INSERT INTO usuarios (permissoes, nome, email, senha, ativo) VALUES(?, ?, ?, ?, TRUE);"
	_, err = DB.Exec(insert, novoUsuario.Permissoes, novoUsuario.Nome, novoUsuario.Email, novoUsuario.Senha)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir usuário no banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Usuário cadastrado com sucesso!"})
}

func AtualizarUsuario(c *gin.Context) {
	var u models.Usuario
	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário. Tente novamente."})
		return
	}

	valido, erroUsu := u.EValido()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": erroUsu})
		return
	}

	admSenha := utils.CriptografarSenha(u.Senha)
	u.Senha = ""

	// Ver se senha do administrador está correta
	rows := DB.QueryRow(`SELECT cod_usu FROM usuarios WHERE senha = ? 
		AND permissoes = 'Administrador' OR permissoes = 'ADM';`, admSenha)

	var codUsuRows int
	err = rows.Scan(&codUsuRows)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Sua senha está incorreta, ou você não tem permissões para editar este usuário."})
		return
	}

	// Verificar se usuário a ser alterado existe
	rows = DB.QueryRow("SELECT cod_usu FROM usuarios WHERE cod_usu = ?;", u.CodUsuario)

	err = rows.Scan(&codUsuRows)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuário não existe, ou não pode ser atualizado."})
		return
	}

	update := "UPDATE usuarios SET permissoes = ?, nome = ?, email = ? WHERE cod_usu = ?;"
	_, err = DB.Exec(update, u.Permissoes, u.Nome, u.Email, u.CodUsuario)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário. Tente novamente."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso!"})
}

func DeletarUsuario(c *gin.Context) {
	codUsu := c.Param("codUsu")

	rows := DB.QueryRow("SELECT cod_usu FROM usuarios WHERE cod_usu = ?;", codUsu)

	var codUsuRows int
	err := rows.Scan(&codUsuRows)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuário não existe, ou não pode ser deletado."})
		return
	}

	update := "UPDATE usuarios SET ativo = FALSE WHERE cod_usu = ?;"
	_, err = DB.Exec(update, codUsu)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desativar usuário."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuário desativado com sucesso!"})
}
