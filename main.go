package main

import (
	"fmt"

	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
	"github.com/Davieas/Industrial-GolangCLP/tagDB"
)

func main() {


	tagDB.TagWindow()
	fmt.Print("Create Window")
	communication.Tcp()
	control.TcpVerification()
}	

