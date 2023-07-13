package main

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
)

func main() {

	var connectionType string

	// Solicitar ao usuário para escolher o tipo de conexão
	fmt.Print("Escolha o tipo de conexão TCP: ")
	fmt.Scanln(&connectionType)

	communication.Tcp()

	control.RunCLP(connectionType)

}
