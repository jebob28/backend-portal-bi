package rotasapi

import (
	dataops "api-bi/cmd/db"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Cadastrar um novo setor
// @Description  Endpoint para cadastrar um novo setor
// @Tags         Setor
// @Param        nome  query  string  true  "Nome do setor a ser cadastrado"
// @Router       /cadastrosetor [post]
func CadastrarNovoSetor(r *gin.Engine) {
	r.POST("/api/bi/v1/cadastrosetor", func(ctx *gin.Context) {
		nome := ctx.Query("nome")
		cadastro, err := dataops.CadastrarSetor(nome)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{"cadastro": cadastro})

	})
}

// @Summary      Cadastrar um novo usuario
// @Description  Endpoint para cadastrar um novo usuario
// @Tags         Usuario
// @Param        nome  query  string  true  "Nome do usuario a ser cadastrado"
// @Param        login  query  string  true  "login do usuario"
// @Param        id_setor  query  int  true  "id do setor"
// @Param        permisao  query  int  true  "permisao"
// @Param        email  query  string  true  "email do usuario"
// @Param        senha  query  string  true  "senha do usuario"
// @Router       /cadastrousuario [post]
func CadastrarNovoUsuario(r *gin.Engine) {
	r.POST("/api/bi/v1/cadastrousuario", func(ctx *gin.Context) {
		// Recuperar parâmetros da requisição
		nome := ctx.Query("nome")
		login := ctx.Query("login")
		idSetor := ctx.Query("id_setor")
		permisao := ctx.Query("permisao")
		email := ctx.Query("email")
		senha := ctx.Query("senha")

		// Validar parâmetros obrigatórios
		if nome == "" || login == "" || email == "" || senha == "" || idSetor == "" || permisao == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Todos os campos são obrigatórios",
			})
			return
		}

		// Converter valores para inteiros
		idSetorInt, err := strconv.Atoi(idSetor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id_setor deve ser um número válido",
			})
			return
		}

		permisaoInt, err := strconv.Atoi(permisao)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "permisao deve ser um número válido",
			})
			return
		}

		// Chamar a função de criação do usuário
		cadastro, err := dataops.CadastrarUsuario(nome, login, email, senha, idSetorInt, permisaoInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Responder com sucesso
		ctx.JSON(http.StatusOK, gin.H{"message": cadastro})
	})
}

// @Summary      Cadastrar link do Bi
// @Description  Endpoint para cadastrar link do BI
// @Tags         BI
// @Param        link  query  string  true  "link do bi"
// @Param        id_setor  query  int  true  "id do setor"
// @Router       /cadastrarbi [post]
func CadastrarBi(r *gin.Engine) {
	r.POST("/api/bi/v1/cadastrarbi", func(ctx *gin.Context) {
		link := ctx.Query("link")
		id_setor := ctx.Query("id_setor")
		idint, err := strconv.Atoi(id_setor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id dever ser um número",
			})
		}
		cadastro, err := dataops.SalvarBi(link, idint)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{"cadastro": cadastro})

	})
}

// @Summary      Atualizar Setor
// @Description  Endpoint para atualizar Setor
// @Tags         Setor
// @Param         id_setor query  int  true  "id do setor responsavel"
// @Param         nome_setor  query  string  true  "novo nome do setor"
// @Router       /atualizasetor [put]
func AtualizarSetor(r *gin.Engine) {
	r.PUT("/api/bi/v1/atualizasetor", func(ctx *gin.Context) {
		id := ctx.Query("id_setor")
		nome := ctx.Query("nome_setor")
		idint, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id dever ser um número",
			})
		}
		updatenosetor, err := dataops.UpdateSetor(idint, nome)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{"atualizacao": updatenosetor})
	})

}

// @Summary      Atualizar usuario
// @Description  Endpoint para atualizar usuario
// @Tags          Usuario
// @Param         usuario query  string  true  "id do setor responsavel"
// @Param         iduser  query  int  true  "novo nome do setor"
// @Param         idsetor  query  int  true  "novo nome do setor"
// @Router       /atualizarusuario [put]
func AtualizarUsuario(r *gin.Engine) {
	r.PUT("/api/bi/v1/atualizarusuario", func(ctx *gin.Context) {
		nome := ctx.Query("usuario")
		id_usuario := ctx.Query("iduser")
		id_setor := ctx.Query("idsetor")
		idusuario, err := strconv.Atoi(id_usuario)
		idsetor, err1 := strconv.Atoi(id_setor)
		if err != nil || err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "iduser e idsetor devem ser números válidos",
			})
			return
		}
		updatenousuario, err := dataops.AtualizarUser(nome, idsetor, idusuario)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"atualizacao": updatenousuario})

	})

}

// @Summary      Atualizar link do BI
// @Description  Endpoint para atualizar Bi
// @Tags          BI
// @Param         idsetor query  int  true  "id do seto responsavel"
// @Param         idbi  query  int  true  "id do bi a ser atualizado "
// @Param         link  query  string  true  "novo link do bi"
// @Router       /atualizarlinkbi [put]
func AtualizarLinkDoBi(r *gin.Engine) {
	r.PUT("/api/bi/v1/atualizarlinkbi", func(ctx *gin.Context) {

		id_setor := ctx.Query("idsetor")
		id_bi := ctx.Query("idbi")
		link := ctx.Query("link")
		idsetor, err := strconv.Atoi(id_setor)
		idbi, err1 := strconv.Atoi(id_bi)
		if err != nil || err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "iduser e idsetor devem ser números válidos",
			})
			return
		}
		updatenousuario, err := dataops.AtualizarBi(idsetor, idbi, link)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"atualizacao": updatenousuario})

	})

}

// @Summary      Deletar Setor
// @Description  Endpoint para Deletar Setor
// @Tags          Setor
// @Param         idsetor query  int  true  "id do seto responsavel"
// @Router       /deletesetor [delete]
func Deletarosetor(r *gin.Engine) {
	r.DELETE("/api/bi/v1/deletesetor", func(ctx *gin.Context) {
		id_setor := ctx.Query("idsetor")
		idsetor, err := strconv.Atoi(id_setor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "idsetor deve ser um número válido",
			})
		}
		deletesetor, err := dataops.DeletarSetor(idsetor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"deletar": deletesetor})

	})
}

// @Summary      Deletar Usuario
// @Description  Endpoint para Deletar Usuario
// @Tags          Usuario
// @Param         id query  int  true  "id do seto responsavel"
// @Router       /deletaruser [delete]
func DeletUser(r *gin.Engine) {
	r.DELETE("/api/bi/v1/deletaruser", func(ctx *gin.Context) {
		id_user := ctx.Query("id")
		id, err := strconv.Atoi(id_user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id deve ser um número válido",
			})
		}

		deleteuser, err := dataops.DeletarUsuario(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"deletar": deleteuser})
	})
}

// @Summary      Deletar link do BI
// @Description  Endpoint para Deletar Link do Bi
// @Tags          BI
// @Param         idlink query  int  true  "id do seto responsavel"
// @Router       /deletarlink [delete]
func DeletLinkBi(r *gin.Engine) {

	r.DELETE("/api/bi/v1/deletarlink", func(ctx *gin.Context) {
		id_link := ctx.Query("idlink")
		idlink, err := strconv.Atoi(id_link)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "idlink deve ser um número válido",
			})
		}
		deletelink, err := dataops.DeletarPortalBi(idlink)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"deletar": deletelink})
	})

}

// @Summary      Selecionar o setor
// @Description  Endpoint para GET Setor
// @Tags          Setor
// @Router       /setor [get]
func Setor(r *gin.Engine) {
	r.GET("/api/bi/v1/setor", func(ctx *gin.Context) {
		setor, err := dataops.SelectSetor()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"setor": setor})
	})

}

// @Summary      Selecionar o usuarios
// @Description  Endpoint para GET usuarios
// @Tags          Usuario
// @Router       /usuarios [get]
func Usuario(r *gin.Engine) {
	r.GET("/api/bi/v1/usuarios", func(ctx *gin.Context) {
		usuario, err := dataops.SelectUser()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{"usuario": usuario})
	})

}

// @Summary      Selecionar o link do BI
// @Description  Endpoint para GET usuarios
// @Tags          BI
// @Param        idsetor query string true  "id do setor"
// @Router       /links [get]
func LinkBI(r *gin.Engine) {
	r.GET("/api/bi/v1/links", func(ctx *gin.Context) {
		// Obtém o parâmetro da query string
		id := ctx.Query("idsetor")

		// Verifica se o parâmetro foi passado
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Parâmetro 'idsetor' é obrigatório",
			})
			return
		}

		// Converte o parâmetro para inteiro
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Parâmetro 'idsetor' deve ser um número válido",
			})
			return
		}

		// Chama a função SelectBi
		link, err := dataops.SelectBi(id_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Erro ao buscar o link BI",
				"detail": err.Error(),
			})
			return
		}

		// Retorna o resultado com sucesso
		ctx.JSON(http.StatusOK, link)
	})
}
