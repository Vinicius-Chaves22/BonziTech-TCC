// security.go inclui todos os middlewares de segurança
// da aplicação
package security

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	models "github.com/vidacalura/BonziTech-TCC/internal/models"
)

func ValidacaoRequest(c *gin.Context) {
	if c.Request.URL.String() == "/api/auth/login" ||
		c.Request.URL.String() == "/api/ping" {
		c.Next()
		return
	}

	// TODO: validar se há sysKey na request

	// TODO: validar se request contém chave do Client (Node.js)

	codSessaoStr := c.Request.Header["Codsessao"]
	if codSessaoStr == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Você precisa estar logado para ter acesso ao sistema"})
		c.Abort()
		return
	}

	codSessao, err := strconv.Atoi(codSessaoStr[0])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Código de sessão inválido"})
		c.Abort()
		return
	}

	// TODO: passar syskey
	// Validar permissões de usuário
	valuesSessao := map[string]int{"codSessao": codSessao}
	jsonValue, _ := json.Marshal(valuesSessao)

	respAuth, err := http.Post(
		"http://"+os.Getenv("dominio")+"/api/auth/usuario", "Application/JSON",
		bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao conectar com o servidor. Tente novamente mais tarde."})
		c.Abort()
		return
	}

	defer respAuth.Body.Close()
	resBody, err := ioutil.ReadAll(respAuth.Body)

	var u models.UsuarioResponse
	if err := json.Unmarshal(resBody, &u); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao conectar com o servidor. Tente novamente mais tarde."})
		c.Abort()
		return
	}

	if u.Error != "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": u.Error})
		c.Abort()
		return
	}

	// TODO: Validar se tipo de request bate com permissões de usuário

	c.Next()
}
