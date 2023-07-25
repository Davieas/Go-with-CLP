package main

import (
	_"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
)

func main() {


	communication.Tcp()

	control.RunCLP()

}

