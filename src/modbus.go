package modbus

import (

"github.com/robinson/gos7"
"github.com/Davieas/Industrial-GolangCLP/config/"
)





func connectionTCP(){

conf := config.TcpClient()
cfgRack := config.GetRackLocation()
cfgSlot := config.GetSlotNumber()


	conn := gos7.NewClient(conf.Client.IP, cfgRack.GetRackLocation(), cfgSlot.GetSlotNumber())
	err := conn.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Disconnect()

}
