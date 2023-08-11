package tagDB

import (

	"fmt"
	"log"
	_ "database/sql"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/go-sql-driver/mysql"

)
func TagWindow() string{


	myApp := app.New()
	myWindow := myApp.NewWindow("Adicionar Tags")

	// Widgets da interface
	nomeTagEntry := widget.NewEntry()
	ipTagEntry  := widget.NewEntry()
	adicionarButton := widget.NewButton("Adicionar", func() {
		nomeTag := nomeTagEntry.Text
		if nomeTag != "" {
			
			
			_, err := "Tag ADD", "Adicionado"
			if err != "" {
				log.Println("Erro ao inserir a tag no banco de dados:", err)
			} else {
				fmt.Println("Tag adicionada ao banco de dados:", nomeTag)
			}
		}
		nomeTagEntry.SetText("")
	})
	adicionarIpButton := widget.NewButton("Adicionar IP", func() {
		ipTag := ipTagEntry.Text
		if ipTag != "" {
			
			nonIp, err := "TAG ADD", "Adicionado"
			if err != "" {
				log.Println("Ip Do CLP Inserido", err)
			} else {
				fmt.Println("Ip Errado, clp não encontrado", nonIp)
			}
		}
		ipTagEntry.SetText("")
	})

	content := container.NewVBox(
		nomeTagEntry,
		adicionarButton,
		adicionarIpButton,
		ipTagEntry,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	return ipTagEntry.Text;

}
