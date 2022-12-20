package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/payments"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := payments.NewGerencianet(credentials)

	params := map[string]string {
		"idPagamento" : "1",
		// "status" : "REALIZADO", // EM_PROCESSAMENTO, AGENDADO, LIQUIDADO, CANCELADO, NAO_REALIZADO 
	}

	res, err := gn.PayDetailPayment(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}