package control

import (
	"bufio"
	"fmt"
	"os"
)

func Authentication() {
	userCrendential := bufio.NewReader(os.Stdin)
	userName, err := userCrendential.ReadString('\n')
	if err != nil {
	  panic("username don´t exist")
	}
	userPassword := bufio.NewReader(os.Stdin)
    pass, err := userPassword.ReadString('\n')
	
	if err != nil {
	  fmt.Println("\r\nError" + err.Error())
	  os.Exit(1)
	}

	if VerifyUser(userName, pass) {
		fmt.Println("Autenticação bem-sucedida")
	} else {
		fmt.Println("Credenciais inválidas")
	}
}

func autorizantion(userRole, resource string) bool {

	//TODO: Implement authorization logic here
	roleResourcePermissions := map[string][]string{
		"admin":     {"editar_params", "excluir_usuarios", "adicionar_recurso"},
		"usuario":   {"visualizar_ip", "visualizar_tags"},
		"developer": {"editar_ip", "visualizar_ip", "visualizar_tags"},
		// Adicione mais papéis e recursos conforme necessário
		// Add more resources....... if you need
	}
	permissions, ok := roleResourcePermissions[userRole]
	if !ok {
		// Se o papel do usuário não estiver no mapeamento, ele não tem permissão
		// if resource, is not in map, he don't have permissions to functions
		return false
	}
	for _, permittedResource := range permissions {
		if permittedResource == resource {
			// O usuário tem permissão para acessar o recurso
			//User have permission to resource
			return true
		}
	}
	return false

}

