package main

import (
	"clinical/api/config/database"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func startServer(ctx *gin.Engine) error {
	log.Printf("Iniciando servidor em http://127.0.0.1:8080 ...")

	//Tratamento de erros para iniciar o servidor
	if err := ctx.Run(":8080"); err != nil {
		log.Printf("Erro ao iniciar o servidor: %v", err)

		return fmt.Errorf("Não foi possivel iniciar o servidor: %w", err)
	}

	return nil
}

func main() {

	//Carregas as variáveis de ambiente do arquivo .env e verifica se ocorreu algum erro
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Erro ao carregar variaveis de ambiente: %v", &err)
	}

	//Conecta ao banco de dados e verifica se ocorreu algum erro
	if _, err := database.Connect(); err != nil{
		log.Fatal(err)
	}

	router := gin.Default()

	//Configuraçãoes de CORS para requisições
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	//Imprime o erro e encerra caso não inicie corretamente o servidor
	if err := startServer(router); err != nil {
		log.Fatal(err)
	}
}
