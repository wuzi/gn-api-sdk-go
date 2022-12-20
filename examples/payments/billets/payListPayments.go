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
		"dataInicio" : "2022-01-01",
		"dataFim" : "2024-12-31",
	}

	res, err := gn.PayListPayments(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}