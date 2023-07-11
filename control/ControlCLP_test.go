package control_test

import (
	"testing"

	"github.com/Davieas/Industrial-GolangCLP/control"
)


func TestControl(t *testing.T){

	if (control.RunCLP("Tcp") == "" ){

		t.Error("Erro encontrado")

	}

}