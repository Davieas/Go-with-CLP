package control

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"log"
)

func RunCLP(connectionType string) string {

	switch connectionType {
	case "Tcp":
		i := ""
		fmt.Println("Digite o Ip do Respectivo CLP:")
		fmt.Scanf(i)
	default:
		fmt.Println("Tipo de conexão inválido.")
		break
	}

	return connectionType

}

func TcpVerification() {

	if err := communication.CheckTcpCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via Ethernet: %v", err)
	}

}