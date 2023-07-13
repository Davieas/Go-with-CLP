package communication_test

import (
	"testing"

	"github.com/Davieas/Industrial-GolangCLP/communication"
)


func TestTcp(t *testing.T){

	if ( communication.Tcp() != ""){

		t.Error("ERRO IN COMMUNICATION REVIEW CODE")

	}
}