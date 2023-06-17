package communication

import (
	"log"
	"fmt"
	"github.com/karalabe/hid"
	"github.com/google/gousb"
)

func USBSerial() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// Verifica se o dispositivo corresponde ao CLP
		return desc.Vendor == gousb.ID(0x1234) && desc.Product == gousb.ID(0x5678)
	})
	if err != nil {
		log.Fatalf("Erro ao abrir o dispositivo: %v", err)
	}

	// Verifica se foi encontrado um dispositivo
	if len(devs) == 0 {
		log.Fatal("Dispositivo CLP não encontrado")
	}

	devices := hid.Enumerate(0, 0)

	//Numero de portas conectadas ao computador(Portas HID)
	for _, device := range devices {
		fmt.Println("Porta serial:", device.Path)

		err := portProcessment(device.Path)
		if err != nil {
			log.Fatalf("Erro ao processar a porta: %v", err)
		}
	}
}

func CheckUSBSerialCommunication() error {
	
	return nil
}


