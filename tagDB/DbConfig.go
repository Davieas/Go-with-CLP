package tagDB

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUsername = "Dinamic-Name"
	DBPassword = "1200"
	DBName     = "plcDB"
)

func TagsPLC() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", DBUsername, DBPassword, DBName))
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Crie a tabela no banco de dados para armazenar as tags (se ainda não existir)
  	criarTabela := `
		CREATE TABLE IF NOT EXISTS tags (
			id INT AUTO_INCREMENT PRIMARY KEY,
			nome VARCHAR(255) NOT NULL
		);
	`	

	
	_, err = db.Exec(criarTabela)

	if err != nil {
		log.Fatal("Erro ao criar a tabela no banco de dados:", err)
		os.Exit(256)
	}

	


}