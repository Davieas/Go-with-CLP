package communication

import (

	"fmt"

)

func portProcessment(port string) error {

		fmt.Println("Porta USB do CLP:", port)

		return nil

}

func TurnOnLamp() {
	// Aqui você adiciona a lógica para ligar a lâmpada no CLP
	fmt.Println("Ligando a lâmpada...")
}

func TurnOffLamp() {
	// Aqui você adiciona a lógica para desligar a lâmpada no CLP
	fmt.Println("Desligando a lâmpada...")
}