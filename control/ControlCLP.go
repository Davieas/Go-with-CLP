package control

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"log"
)

func RunCLP(connectionType string){


	switch connectionType {
	case "USB":
		communication.USBSerial()
		usbSerialVerification()
		
	case "Tcp":
		communication.Tcp()
		TcpVerification()
	default:
		fmt.Println("Tipo de conexão inválido.")
	}

	fmt.Println("Controlando o CLP...")


}
func usbSerialVerification() {
	// Verifica a comunicação via USB Serial
	if err := communication.CheckUSBSerialCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via USB Serial: %v", err)
	}

	
}

func TcpVerification(){

	if err := communication.CheckTcpCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via Ethernet: %v", err)
	}


}
