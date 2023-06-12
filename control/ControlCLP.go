package control

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
)

func RunCLP(){

	fmt.Println("Controlando o CLP...")
	communication.TurnOnLamp()
	communication.TurnOffLamp()

}

