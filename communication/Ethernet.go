package communication


import (

	"fmt"
	"log"
	"github.com/goburrow/modbus"
	"time"
	

)


func Tcp(i string) string{
	
	tcpIpHandler := modbus.NewTCPClientHandler("192.168.0.1:8080")
	handler := tcpIpHandler  // IP e porta do CLP
	handler.Timeout = 2 * time.Second                      // Tempo limite para a comunicação
	err := handler.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao CLP via Tcp: %v", err)
	}
	defer handler.Close()
	
	fmt.Print("Tcp Connection closed!!!!")

	client := modbus.NewClient(handler)

	address := uint16(0x0001) // Endereço do registrador para controlar a lâmpada
	coilStatus := uint16(0x1)       // Ligar a lâmpada

    _ ,err = client.WriteSingleCoil(address, coilStatus)
	if err != nil {
		log.Fatalf("Erro ao ligar a lâmpada: %v", err)
	}
	fmt.Println("Lâmpada ligada!")

	// Aguardar um tempo para a lâmpada ficar ligada
	time.Sleep(2 * time.Second)

	// Desligar a lâmpada
	coilStatus = uint16(0x1)
	_,err = client.WriteSingleCoil(address, coilStatus)
	if err != nil {
		log.Fatalf("Erro ao desligar a lâmpada: %v", err)
	}
	fmt.Println("Lâmpada desligada!")

	return "100"
}

func CheckTcpCommunication() error {
	return nil
}

