package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

type Transaction struct {
	//ID                 uuid.UUID
	TransactionValue   float64
	ProductDescription string
	CardNumber         float64
	NameInCard         string
	CardExpirationDate time.Time
	Cvv                int8
}

func PGConnection() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("erro ao tentar ler env variable")
	}

	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PSW := os.Getenv("PSW")
	DBNAME := os.Getenv("DBNAME")
	PORT := os.Getenv("PORT")

	sqlDB, err := sql.Open("pgx", "host="+HOST+" user="+USER+" password="+PSW+" dbname="+DBNAME+" port="+PORT+" sslmode=disable")

	if err != nil {
		fmt.Println("Erro ao tentar abrir uma conexao com o banco de dados")
	}

	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados")

	} else {
		t := &Transaction{
			//ID:                 uuid.New(),
			TransactionValue:   65456.55,
			ProductDescription: "teste",
			CardNumber:         1234567890987654,
			NameInCard:         "J T S",
			CardExpirationDate: time.Now(),
			Cvv:                123,
		}
		fmt.Println("Iniciando criacao de transacao")
		if err := GormDB.Create(&t).Error; err != nil {
			fmt.Println("Ocorreu um erro ao tentar salvar transacao")
		} else {
			fmt.Println("Dado criado com sucesso")
		}

		fmt.Println("Iniciando a consulta de todas as records")
		var x []Transaction
		if err := GormDB.Raw("SELECT * FROM transactions").Scan(&x).Error; err != nil {
			fmt.Printf("ocorreu um erro ao tentar consultar")
		} else {
			print(x)
		}
	}
}
