package main

import (
<<<<<<< HEAD
	"fmt"

=======
	_"fmt"
>>>>>>> origin/master
	"github.com/Davieas/Industrial-GolangCLP/communication"
	"github.com/Davieas/Industrial-GolangCLP/control"
	"github.com/Davieas/Industrial-GolangCLP/tagDB"
)

func main() {


<<<<<<< HEAD
	tagDB.TagWindow()
	fmt.Print("Tela Criada")
	communication.Tcp()
	control.RunCLP()
}	

=======
	communication.Tcp()

	control.RunCLP()

}

>>>>>>> origin/master
