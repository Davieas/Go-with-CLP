package communication

import (

	"fmt"
	"net"

)

func portProcessment(port string) error {

		fmt.Println("Porta USB do CLP:", port)

		return nil

}
func GetTCPIP() (string, error) {
	
	IP := ""
	fmt.Println("Digite o Ip do seu CLP: ")
	fmt.Scanln(IP)

	// Verifique se o endereço IP do CLP é válido
	if ip := net.ParseIP(IP); ip == nil {
		return "", fmt.Errorf("Endereço IP inválido do CLP: %s", IP)
	}

	return IP, nil
}
