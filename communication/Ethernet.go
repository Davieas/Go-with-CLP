package communication


import (

	"fmt"
	"log"
	"github.com/goburrow/modbus"
	"github.com/tarm/serial"
	"time"
	"github.com/Davieas/Industrial-GolangCLP/tagDB"
	

)


func Tcp() string{

	c := &serial.Config{Name: "COM5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
			log.Fatal(err)
	}

	
	tcpIpHandler := modbus.NewTCPClientHandler(tagDB.TagWindow())
	handler := tcpIpHandler  // IP e porta do CLP
	handler.Timeout = 2 * time.Second                      // Tempo limite para a comunicação
	err = handler.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao CLP via Tcp: %v", err)
	}
	defer handler.Close()
	
	fmt.Print("Tcp Connection closed!!!!")

	client := modbus.NewClient(handler)

	buf := make([]byte, 128)
      _, err = s.Read(buf)
      if err != nil {
              log.Fatal(err)
      }

	

	address := uint16(0x0001) // Endereço do registrador para controlar a lâmpada
	coilStatus := uint16(0x1)       // Ligar a lâmpada

    _ ,err = client.WriteSingleCoil(address, coilStatus)
	if err != nil {
		log.Fatalf("Erro ao ligar a lâmpada: %v", err)
	}
	fmt.Println("Lâmpada ligada!")

	time.Sleep(2 * time.Second)
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