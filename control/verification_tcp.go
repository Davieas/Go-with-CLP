package control

import (
	_"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"log"
)


func TcpVerification() {

	if err := communication.CheckTcpCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via Ethernet: %v", err)
	}

}