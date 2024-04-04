package main

import (
	"fmt"

	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
	"github.com/Davieas/Industrial-GolangCLP/tagDB"
)

func main() {

control.Authentication()
control.VerifyUser()
communication.TcpConnection()
tagDB.TagWindow()
// Printing the current tags in the database
fmt.Print("Succefuly, connected CLP!!")
}	

