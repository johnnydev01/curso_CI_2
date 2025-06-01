package database

import (
	"log"
	"fmt"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")

	if endereco == "" || usuario == "" || senha == "" || nomeBanco == "" || portaBanco == "" || sslMode == ""{
		panic("Uma ou mais variáveis de ambiente do banco de dados estão vazias")
	}

	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		endereco, usuario, senha, nomeBanco, portaBanco, sslMode,
	)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
