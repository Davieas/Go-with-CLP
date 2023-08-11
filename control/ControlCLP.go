package control

import (
	_"fmt"
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"log"
)

func RunCLP()  {

<<<<<<< HEAD
//i := ""
//fmt.Println("Digite o Ip do Respectivo CLP:")
//fmt.Scanf(i)
=======
		i := ""
		fmt.Println("Digite o Ip do Respectivo CLP:")
		fmt.Scanf(i)
>>>>>>> origin/master


}

func TcpVerification() {

	if err := communication.CheckTcpCommunication(); err != nil {
		log.Fatalf("Erro de comunicação via Ethernet: %v", err)
	}

}