package main

import (
	"fmt"
	_"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
)

func main() {

	var connectionType string

	// Solicitar ao usuário para escolher o tipo de conexão
	fmt.Print("Escolha o tipo de conexão TCP: ")
	fmt.Scanln(&connectionType)

	control.RunCLP(connectionType)

}
