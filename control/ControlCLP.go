package control

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"log"
)

func RunCLP()  {

		i := ""
		fmt.Println("Digite o Ip do Respectivo CLP:")
		fmt.Scanf(i)


}

func TcpVerification() {

	if err := communication.CheckTcpCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via Ethernet: %v", err)
	}

}