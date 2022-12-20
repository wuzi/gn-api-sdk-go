package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "60740262386",
		"phone_number": "51944916523",
		"email": "gorb.oldbuck@gerencianet.com.br",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"banking_billet": map[string]interface{} {
				"expire_at": "2022-12-12",
				"customer": customer,
			},
		},
		"items": []map[string]interface{}{
			{
				"name": "Product 1",
				"value": 500,
				"amount": 2,
			},
		},
	}

	res, err := gn.CreateOneStepCharge(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}