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
		"codBarras" : "36400000000000000000000000000000000000000000000",
	}

	body := map[string]interface{} {
		"valor" : 500,
		"dataPagamento" : "2023-01-01",
		"descricao" : "Payment of the test billet",
	}

	res, err := gn.PayRequestBarCode(params, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}