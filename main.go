package main

import (
	rotasapi "api-bi/cmd/routes"
	_ "api-bi/docs"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func GerarRoutesJson(router *gin.Engine) {
	routes := router.Routes()

	type Route struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	// Cria a lista de rotas
	var routesJson []Route
	for _, route := range routes {
		routesJson = append(routesJson, Route{
			Method: route.Method,
			Path:   route.Path,
		})
	}

	// Cria o arquivo routes.json
	file, err := os.Create("routes.json")
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo routes.json: %v\n", err)
		return
	}
	defer file.Close()

	// Codifica os dados em JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(routesJson); err != nil {
		fmt.Printf("Erro ao codificar as rotas em JSON: %v\n", err)
		return
	}

	// Mensagem de sucesso
	fmt.Println("Arquivo routes.json gerado com sucesso!")
}

// @title           Swagger Examplo API Do BI
// @version         1.0
// @host      localhost:8080
// @BasePath  /api/bi/v1/
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//Rotas
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"}, // Origens permitidas
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                 // Métodos permitidos
		AllowHeaders:     []string{"Content-Type", "Authorization"},                // Cabeçalhos permitidos
		ExposeHeaders:    []string{"Content-Length"},                               // Cabeçalhos expostos
		AllowCredentials: true,                                                     // Permitir credenciais
	}))
	//route cadastrar setor
	rotasapi.CadastrarNovoSetor(r)
	//route cadastrar funcionario
	rotasapi.CadastrarNovoUsuario(r)
	// route cadastrar bi
	rotasapi.CadastrarBi(r)
	// route atualizar setor
	rotasapi.AtualizarSetor(r)
	//route atualizar usuario
	rotasapi.AtualizarUsuario(r)
	//route atualizar link do bi
	rotasapi.AtualizarLinkDoBi(r)
	// route deletar setor
	rotasapi.Deletarosetor(r)
	// route deletar usuario
	rotasapi.DeletUser(r)
	// route deletar bi
	rotasapi.DeletLinkBi(r)

	// route buscar setor
	rotasapi.Setor(r)
	// route buscar usuario
	rotasapi.Usuario(r)
	// route buscar bi
	rotasapi.LinkBI(r)

	GerarRoutesJson(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
