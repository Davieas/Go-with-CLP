package main

import (
	"fmt"
	"time"

	"github.com/robinson/gos7"
)

func main() {
	// Cria uma conexão com o CLP
	conn := gos7.NewClient(192.168.0.100, 0x01, 0x10)
	err := conn.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Disconnect()

	// Liga o motor
	err = conn.WriteBySymbol("DB1.DBX0.0", 1)
	if err != nil {
		panic(err)
	}

	// Espera 5 segundos antes de ligar a bomba
	time.Sleep(5 * time.Second)

	// Liga a bomba
	err = conn.WriteBySymbol("DB1.DBX2.0", 1)
	if err != nil {
		panic(err)
	}

	// Lê o valor do sensor de nível
	sensorValue, err := conn.ReadFloatBySymbol("DB1.DBD4")
	if err != nil {
		panic(err)
	}

	// Verifica se o nível está alto
	if sensorValue >= 100.0 {
		fmt.Println("Nível alto alcançado!")
	} else {
		fmt.Println("Nível baixo ainda.")
	}
}
