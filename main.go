package main

import (
	"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"


)

func main() {
	fmt.Println("CLP RUNNN!!!")
	communication.USBSerial()
	control.RunCLP()
	
}
