package control

import (
		"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func VerifyUser(username, password string) bool {
	// Conectar ao banco de dados
	//Connect on Database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		// Tratar o erro de conexão com o banco de dados
		// Handle database connection error
		return false
	}
	defer db.Close()

	// Consulta para verificar se as credenciais são válidas
	//Query on credentials for validation

	query := "SELECT COUNT(*) FROM usuarios WHERE usuario = ? AND senha = ?"
	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		// Tratar o erro ao executar a consulta
		return false
	}
	// Se a contagem for maior que 0, as credenciais são válidas
	return count > 0
}
